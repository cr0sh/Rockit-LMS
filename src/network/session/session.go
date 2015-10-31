package session

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"net"
	"network/packet"
	"network/protocol"
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
	mtuSize         uint16
	connectionState byte
	sendSeqNum      uint32
	channelIndex    [32]int32
	asyncTicker     *time.Ticker
	splitPackets    map[int16]map[int32][]byte
	needPong        int64
	gotPong         bool
}

const (
	unconnected byte = 0
	connecting1 byte = 1
	connecting2 byte = 2
	connected   byte = 3
)

//Handle Handles packets from RecvStream channel
func (session *Session) Handle() {
	session.splitPackets = make(map[int16]map[int32][]byte)
	session.asyncTicker = time.NewTicker(time.Second * 7)
	go session.asyncProcess()
	for pk := range session.RecvStream {
		switch pk.Head {
		case 0x01: //UNCONNECTED_PING
			fallthrough
		case 0x02: //UNCONNECTED_PING_OPEN_CONNECTION
			continue
		default:
			head := pk.Head
			buf := pk.GetBytes()
			if !session.handlePacket(pk) {
				fmt.Print("Unexpected packet header: 0x", hex.EncodeToString([]byte{head}), "\n"+hex.Dump(buf))
			}
		}
	}
}

func (session *Session) handlePacket(pk packet.Packet) bool {
	if session.connectionState > 1 && pk.Head >= 0x80 && pk.Head < 0x90 {
		var dp packet.DataPacket
		var err error
		if dp, err = packet.NewDataPacket(pk); err != nil {
			fmt.Println("Error while decoding data packet:", err)
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
		fmt.Println("set state to 1")
		return true
	case pk.Head == 0x07:
		pk.Next(16) //Magic
		if _, err := packet.ReadAddress(pk.Buffer); err != nil {
			logging.FromError(err, 0)
			fmt.Print(hex.EncodeToString([]byte{pk.Head}), "\n"+hex.Dump(append([]byte{pk.Head}, pk.Buffer.Bytes()...)))
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
		return true
	case pk.Head == 0xa0: //NACK
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

func (session *Session) asyncProcess() {
	for range session.asyncTicker.C {
		session.gotPong = false
		session.needPong = rand.Int63()
		pk := packet.NewPacket(0x00)
		binary.Write(pk, binary.BigEndian, session.needPong)
		ep := *new(packet.EncapsulatedPacket)
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
		<-time.NewTimer(time.Second * 5).C // timeout
		if !session.gotPong {
			session.Close("Ping timeout")
		}
	}
}

func (session *Session) handleDataPacket(dp packet.DataPacket) bool {
	logging.Debug("Seqnumber " + strconv.Itoa(int(dp.SeqNumber)))
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

func (session *Session) handleEncapsulatedPacket(pk packet.Packet) {
	logging.Debug("Handling DataPacket head 0x" + hex.EncodeToString([]byte{pk.Head}))
	switch pk.Head {
	case 0x00:
		var pingID int64
		if err := binary.Read(pk, binary.BigEndian, &pingID); err != nil {
			return
		}
		pk = packet.NewPacket(0x03)
		binary.Write(pk, binary.BigEndian, pingID)
		ep := *new(packet.EncapsulatedPacket)
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
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
			logging.Verbose("Client", session.Address.String(), "finally connected on Raknet level")
		}
	case 0x15:
		session.Close("client disconnect")
	}
}

func (session *Session) handleSplitPacket(ep packet.EncapsulatedPacket, pk packet.Packet) {
	logging.Debug("SplitID", ep.SplitID, "SplitIndex", ep.SplitIndex, "SplitCount", ep.SplitCount)
	if _, ok := session.splitPackets[ep.SplitID]; !ok {
		session.splitPackets[ep.SplitID] = make(map[int32][]byte)
	}
	if _, ok := session.splitPackets[ep.SplitID][ep.SplitIndex]; !ok {
		session.splitPackets[ep.SplitID][ep.SplitIndex] = pk.GetBytes()
	}
	buffer := new(bytes.Buffer)
	logging.Debug("SC", ep.SplitCount)
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

func (session *Session) sendDataPacket(pk packet.EncapsulatedPacket) {
	dp := *new(packet.DataPacket)
	dp.Buffer = new(bytes.Buffer)
	dp.SeqNumber = session.sendSeqNum
	session.sendSeqNum++
	if session.sendSeqNum == 1<<6 {
		session.sendSeqNum = 0
	}
	dp.Packets = []packet.Packet{packet.Packet{Buffer: bytes.NewBuffer(pk.Bytes()), Head: 0, Address: *new(net.UDPAddr)}}
	session.SendStream <- dp.Encode(0x80)
}

//Close Closes session channels for stopping goroutines
func (session *Session) Close(reason string) {
	pk := packet.NewPacket(0x15)
	ep := *new(packet.EncapsulatedPacket)
	ep.Encapsulate(pk)
	session.sendDataPacket(ep)
	time.Sleep(time.Millisecond * 500) //Wait for 0.5 secs
	logging.Verbose("Session", session.Address.String(), "closed:", reason)
	session.asyncTicker.Stop()
	close(session.RecvStream)
	close(session.SendStream)
}
