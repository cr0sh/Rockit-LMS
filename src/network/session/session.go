package session

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
	"math/rand"
	"net"
	"network/packet"
	"network/protocol"
	"player"
	"strconv"
	"time"
	"util/logging"
)

//Session struct contains packet channels for packet I/O, connection state, and source address.
type Session struct {
	Address         net.UDPAddr
	RecvStream      chan packet.Packet
	SendStream      chan packet.Packet
	ServerID        uint64
	PlayerHandler   player.Handler
	mtuSize         uint16
	connectionState byte
	sendSeqNum      uint32
	channelIndex    [32]uint32
	asyncPingTicker *time.Ticker
	ackTicker       *time.Ticker
	splitPackets    map[uint16]map[uint32][]byte
	needPong        uint64
	gotPong         bool
	recoveryQueue   map[uint32]packet.Packet
	receivedWindow  map[uint32]uint32
	windowBorder    [2]uint32
	lastSeq         uint32
	ackQueue        []uint32
	nackQueue       []uint32
	closed          bool
}

const (
	unconnected  byte = 0
	connecting1  byte = 1
	connecting2  byte = 2
	connected    byte = 3
	recoverysize uint = 128
	windowsize   uint = 32
)

//Handle is a loop for handling packets from RecvStream channel.
func (session *Session) Handle() {
	session.splitPackets = make(map[uint16]map[uint32][]byte)
	session.asyncPingTicker = time.NewTicker(time.Second * 7)
	session.ackTicker = time.NewTicker(time.Millisecond * 200)
	session.recoveryQueue = make(map[uint32]packet.Packet)
	session.receivedWindow = make(map[uint32]uint32)
	session.windowBorder[1] = uint32(windowsize) // [start|end)
	go session.asyncPing()
	go session.acknowledgement()
	for pk := range session.RecvStream {
		if session.closed {
			return
		}
		switch pk.Head {
		case 0x01: //UNCONNECTED_PING
			fallthrough
		case 0x02: //UNCONNECTED_PING_OPEN_CONNECTION
			continue
		default:
			head := pk.Head
			buf := pk.GetBytes()
			if !session.handlePacket(pk) {
				logging.Debug("Unexpected packet header: 0x", hex.EncodeToString([]byte{head}), "\n"+hex.Dump(buf))
			}
		}
	}
}

func (session *Session) handlePacket(pk packet.Packet) bool {
	if session.connectionState > 1 && pk.Head >= 0x80 && pk.Head < 0x90 {
		var dp packet.DataPacket
		var err error
		if dp, err = packet.NewDataPacket(pk); err != nil {
			logging.Debug("Error while decoding data packet:", err)
			return true
		}
		return session.handleDataPacket(dp)
	}
	switch {
	case pk.Head == 0x05:
		pk.Next(16) //Magic
		if len(pk.Buffer.Bytes()) < 18 {
			logging.Error("Error while processing pacekt parse: buffer too short")
			return true
		} else if proto, err := pk.ReadByte(); err == nil && int(proto) != protocol.RaknetProtocol {
			logging.Error("Raknet protocol mismatch: " + strconv.Itoa(int(pk.Buffer.Bytes()[16])) + " != " + strconv.Itoa(protocol.RaknetProtocol))
			return true
		} else if err != nil {
			logging.FromError(err, 0)
			return true
		}
		mtusize := make([]byte, 2)
		if _, err := pk.Read(mtusize); err != nil {
			logging.FromError(err, 0)
			return true
		}
		session.mtuSize = uint16(math.Min(float64(binary.BigEndian.Uint16(mtusize)+18), 1464))
		pk := packet.NewPacket(0x06)
		pk.Buffer.Write(protocol.RaknetMagic)
		binary.Write(pk, binary.BigEndian, session.ServerID)
		pk.WriteByte(0)
		binary.Write(pk, binary.BigEndian, mtusize)
		session.SendStream <- pk
		session.connectionState = connecting1
		logging.Debug("set state to 1")
		return true
	case pk.Head == 0x07:
		pk.Next(16) //Magic
		if _, err := packet.ReadAddress(pk.Buffer); err != nil {
			logging.FromError(err, 0)
			logging.Debug(hex.EncodeToString([]byte{pk.Head}), "\n"+hex.Dump(append([]byte{pk.Head}, pk.Buffer.Bytes()...)))
			return true
		}
		pk = packet.NewPacket(0x08)
		pk.Write(protocol.RaknetMagic)
		binary.Write(pk.Buffer, binary.BigEndian, session.ServerID)
		packet.PutAddress(session.Address, pk.Buffer, 4)
		binary.Write(pk.Buffer, binary.BigEndian, session.mtuSize)
		session.SendStream <- pk
		session.connectionState = connecting2
		logging.Debug("set state to 2")
		return true
	case pk.Head == 0xc0: //ACK
		ack := packet.AcknowledgePacket{bytes.NewBuffer(pk.Bytes()), make([]uint32, 0)}
		ack.Decode()
		for pk := range ack.Packets {
			if _, ok := session.recoveryQueue[uint32(pk)]; ok {
				delete(session.recoveryQueue, uint32(pk))
			}
		}
		return true
	case pk.Head == 0xa0: //NACK
		nack := packet.AcknowledgePacket{bytes.NewBuffer(pk.Bytes()), make([]uint32, 0)}
		nack.Decode()
		for pk := range nack.Packets {
			if rpk, ok := session.recoveryQueue[uint32(pk)]; ok {
				session.SendStream <- rpk
			}
		}
		return true
	case session.connectionState == connecting2 || session.connectionState == connected:
		var dp packet.DataPacket
		var err error
		if dp, err = packet.NewDataPacket(pk); err != nil {
			logging.FromError(err, 0)
			return true
		}
		return session.handleDataPacket(dp)
	}
	return false
}

func (session *Session) handleDataPacket(dp packet.DataPacket) bool {
	logging.Debug("Seqnumber " + strconv.Itoa(int(dp.SeqNumber)))
	session.ackQueue = append(session.ackQueue, dp.SeqNumber)
	for seq := range session.receivedWindow {
		if seq < session.windowBorder[0] {
			delete(session.receivedWindow, seq)
			logging.Debug("recvWindow: clean", seq)
			continue
		}
		break
	}
	if _, ok := session.receivedWindow[dp.SeqNumber]; ok || dp.SeqNumber < session.windowBorder[0] || dp.SeqNumber >= session.windowBorder[1] {
		return true
	}
	diff := dp.SeqNumber - session.lastSeq
	if diff != 1 {
		logging.Debug("Packet loss: diff is", diff)
		for i := session.lastSeq + 1; i < dp.SeqNumber; i++ {
			if _, ok := session.receivedWindow[dp.SeqNumber]; !ok {
				session.nackQueue = append(session.nackQueue, i)
				logging.Debug("Packet loss: requesting", i)
			}
		}
	}
	if diff > 0 {
		session.lastSeq = dp.SeqNumber
		session.windowBorder[0] += diff
		session.windowBorder[1] += diff
	}
	for i, pk := range dp.Packets {
		if dp.EncapsulatedPackets[i].HasSplit {
			logging.Debug("handling split")
			session.handleSplitPacket(dp.EncapsulatedPackets[i], pk)
		} else {
			session.handleEncapsulatedPacket(pk)
		}
	}
	return true
}

func (session *Session) handleSplitPacket(ep packet.EncapsulatedPacket, pk packet.Packet) {
	logging.Debug("Split result: SplitID", ep.SplitID, "SplitIndex", ep.SplitIndex, "SplitCount", ep.SplitCount)
	if ep.SplitCount > 1024 {
		logging.Debug("Oops: invalid packet", hex.Dump(ep.Bytes()))
		session.Close("Bad client")
		return
	}
	if _, ok := session.splitPackets[ep.SplitID]; !ok {
		session.splitPackets[ep.SplitID] = make(map[uint32][]byte)
	}
	if _, ok := session.splitPackets[ep.SplitID][ep.SplitIndex]; !ok {
		session.splitPackets[ep.SplitID][ep.SplitIndex] = pk.GetBytes()
	}
	buffer := new(bytes.Buffer)
	for i := 0; i < int(ep.SplitCount); i++ {
		if buf, ok := session.splitPackets[ep.SplitID][ep.SplitIndex]; !ok {
			logging.Debug("Cannot handle split: need 0 <= n <=", ep.SplitCount-1, "missing:", i)
			break
		} else {
			buffer.Write(buf)
		}
	}
	ppk := *new(packet.Packet)
	var err error
	if ppk.Head, err = buffer.ReadByte(); err != nil {
		return
	}
	ppk.Buffer = bytes.NewBuffer(buffer.Bytes())
	logging.Debug("Handled split")
	session.handleEncapsulatedPacket(ppk)
}

func (session *Session) handleEncapsulatedPacket(pk packet.Packet) {
	logging.Debug("Handling DataPacket head 0x" + hex.EncodeToString([]byte{pk.Head}))
	if pk.Head >= 0x80 && session.connectionState == connected {
		logging.Debug("Forwarding packet to player")
		session.PlayerHandler.HandlePacket(pk)
		return
	}
	switch pk.Head {
	case 0x00:
		var pingID uint64
		if err := binary.Read(pk, binary.BigEndian, &pingID); err != nil {
			return
		}
		pk = packet.NewPacket(0x03)
		binary.Write(pk, binary.BigEndian, pingID)
		ep := *new(packet.EncapsulatedPacket)
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
	case 0x03:
		var pingID uint64
		if err := binary.Read(pk, binary.BigEndian, &pingID); err != nil {
			return
		}
		if pingID == session.needPong {
			session.gotPong = true
			logging.Debug("Got correct pong!")
		}
	case 0x09:
		var cid, sendPing int64
		if err := binary.Read(pk, binary.BigEndian, &cid); err != nil {
			logging.Error(packet.NewError(pk.Buffer, err), 0)
			return
		}
		if err := binary.Read(pk, binary.BigEndian, &sendPing); err != nil {
			logging.Error(packet.NewError(pk.Buffer, err), 0)
			return
		}
		pk = packet.NewPacket(0x10)
		packet.PutAddress(session.Address, pk.Buffer, 4)
		pk.Write([]byte{0, 0})
		if addr, err := net.ResolveUDPAddr("", "127.0.0.1:0"); err == nil {
			packet.PutAddress(*addr, pk.Buffer, 4)
		} else {
			logging.Error(err, 0)
			return
		}
		for i := 0; i < 9; i++ {
			packet.PutAddress(net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 0, Zone: ""}, pk.Buffer, 4)
		}
		binary.Write(pk, binary.BigEndian, sendPing)
		binary.Write(pk, binary.BigEndian, sendPing+1000)
		ep := *new(packet.EncapsulatedPacket)
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
	case 0x13:
		if _, err := packet.ReadAddress(pk.Buffer); err == nil {
			session.connectionState = connected
			logging.Verbose("Client", session.Address.String(), "finally connected on Raknet level")
			session.PlayerHandler = player.Handler{Address: session.Address}
		}
	case 0x15:
		session.Close("client disconnect")
	}
}

func (session *Session) asyncPing() {
	for range session.asyncPingTicker.C {
		session.gotPong = false
		session.needPong = uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
		pk := packet.NewPacket(0x00)
		binary.Write(pk, binary.BigEndian, session.needPong)
		ep := *new(packet.EncapsulatedPacket)
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
		<-session.asyncPingTicker.C
		if !session.gotPong {
			session.Close("Ping timeout")
		}
	}
}

func (session *Session) acknowledgement() {
	for range session.ackTicker.C {
		if len(session.ackQueue) > 0 {
			ack := *new(packet.AcknowledgePacket)
			ack.Packets = session.ackQueue
			ack.Encode()
			pk := packet.NewPacket(0xc0)
			*pk.Buffer = *ack.Buffer
			session.SendStream <- pk
			session.ackQueue = make([]uint32, 0)
		}
		if len(session.nackQueue) > 0 {
			nack := *new(packet.AcknowledgePacket)
			nack.Packets = session.nackQueue
			nack.Encode()
			pk := packet.NewPacket(0xa0)
			*pk.Buffer = *nack.Buffer
			session.SendStream <- pk
			session.nackQueue = make([]uint32, 0)
		}
	}
}

func (session *Session) sendDataPacket(pk packet.EncapsulatedPacket) {
	dp := *new(packet.DataPacket)
	dp.Buffer = new(bytes.Buffer)
	dp.SeqNumber = session.sendSeqNum
	dp.Packets = []packet.Packet{packet.Packet{Buffer: bytes.NewBuffer(pk.Bytes()), Head: 0, Address: *new(net.UDPAddr)}}
	rpk := dp.Encode(0x80)
	if recoverysize > 0 {
		if pk.NeedACK {
			session.recoveryQueue[session.sendSeqNum] = rpk
		}
		tmp := session.sendSeqNum - uint32(recoverysize)
		if tmp < 0 {
			tmp = 1<<6 + tmp
		}
		if _, ok := session.recoveryQueue[tmp]; ok {
			delete(session.recoveryQueue, tmp)
		}
	}
	session.sendSeqNum++
	if session.sendSeqNum == 1<<6 {
		session.sendSeqNum = 0
	}
	if !session.closed {
		session.SendStream <- rpk
	}
}

//Close closes session, related channels, and sends disconnect signal to client.
func (session *Session) Close(reason string) {
	pk := packet.NewPacket(0x15)
	ep := *new(packet.EncapsulatedPacket)
	ep.Encapsulate(pk)
	session.sendDataPacket(ep)
	session.closed = true
	time.Sleep(time.Millisecond * 500) //Wait for 0.5 secs
	logging.Verbose("Session", session.Address.String(), "closed:", reason)
	session.asyncPingTicker.Stop()
	session.ackTicker.Stop()
	close(session.RecvStream)
	close(session.SendStream)
}
