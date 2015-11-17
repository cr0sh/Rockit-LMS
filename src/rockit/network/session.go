package network

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
	"net"
	"rockit/network/packet"
	"rockit/util/logger"
	"runtime"
	"sort"
	"strconv"
	"time"
)

type packetRecovery struct {
	sendTime  time.Time
	seqNumber uint32
	packet    []byte
}

const (
	unconnected   = 0
	connecting1   = 1
	connecting2   = 2
	connected     = 3
	maxSplitSize  = 128
	maxSplitCount = 4
	windowSize    = 2048
)

//Session struct contains packet channels for packet I/O, connection state, and source address.
type Session struct {
	Address             net.UDPAddr
	ServerID            uint64
	SessionID           uint
	RecvStream          chan packet.Packet
	SendStream          chan []byte
	state               byte
	isActive            bool
	lastUpdate          time.Time
	sendSeq             uint32
	lastSeq             uint32
	packetToSend        map[uint32]packetRecovery
	recoveryQueue       map[uint32]packetRecovery
	ackQueue            map[uint32]uint32
	nackQueue           map[uint32]uint32
	receivedWindow      map[uint32]packet.DataPacket
	windowStart         int64
	windowEnd           int64
	lastReliableIndex   int64
	reliableWindowStart uint32
	reliableWindowEnd   uint32
	reliableWindow      map[uint32]packet.EncapsulatedPacket
	messageIndex        uint32
	channelIndex        map[uint16]uint32
	sendQueue           packet.DataPacket
	mtuSize             uint64 //Should be > 36
	splitID             uint16
	splitPackets        map[uint16]map[uint32][]byte
}

//HandleSession is a core loop for session processing
func (session *Session) HandleSession() {
	session.isActive = false
	session.lastUpdate = time.Now()
	session.packetToSend = make(map[uint32]packetRecovery)
	session.recoveryQueue = make(map[uint32]packetRecovery)
	session.ackQueue = make(map[uint32]uint32)
	session.nackQueue = make(map[uint32]uint32)
	session.receivedWindow = make(map[uint32]packet.DataPacket)
	session.channelIndex = make(map[uint16]uint32)
	session.reliableWindow = make(map[uint32]packet.EncapsulatedPacket)
	session.sendQueue = packet.DataPacket{
		Buffer:              new(bytes.Buffer),
		SeqNumber:           0,
		Head:                0x84,
		EncapsulatedPackets: make([]packet.EncapsulatedPacket, 0),
		Packets:             make([][]byte, 0),
	}
	session.splitPackets = make(map[uint16]map[uint32][]byte)
	session.windowStart = -1
	session.windowEnd = windowSize
	session.reliableWindowStart = 0
	session.reliableWindowEnd = windowSize
	update := time.NewTicker(time.Millisecond * 50)
	go session.updateJob(update.C)
	defer func() {
		update.Stop()
		logger.Debug("Stopping session")
	}()
	for pk := range session.RecvStream {
		session.handlePacket(pk)
	}
}

func (session *Session) updateJob(tick <-chan time.Time) {
	for range tick {
		if !session.isActive && session.lastUpdate.Add(time.Second*10).Before(time.Now()) {
			session.disconnect("timeout")
			return
		}
		if len(session.ackQueue) > 0 {
			ack := new(packet.AcknowledgePacket)
			ack.Packets = make([]uint32, len(session.ackQueue))
			for i := range session.ackQueue {
				ack.Packets = append(ack.Packets, i)
			}
			ack.Encode()
			session.SendStream <- append([]byte{0xc0}, ack.Bytes()...)
			session.ackQueue = make(map[uint32]uint32)
		}
		if len(session.nackQueue) > 0 {
			nack := new(packet.AcknowledgePacket)
			nack.Packets = make([]uint32, len(session.nackQueue))
			for i := range session.nackQueue {
				nack.Packets = append(nack.Packets, i)
			}
			nack.Encode()
			session.SendStream <- append([]byte{0xa0}, nack.Bytes()...)
			session.nackQueue = make(map[uint32]uint32)
		}
		if len(session.packetToSend) > 0 {
			limit := 16
			for k, v := range session.packetToSend {
				v.sendTime = time.Now()
				session.recoveryQueue[v.seqNumber] = v
				delete(session.packetToSend, k)
				session.SendStream <- v.packet
				limit--
				if limit <= 0 {
					break
				}
			}
			if len(session.packetToSend) > windowSize {
				session.packetToSend = make(map[uint32]packetRecovery)
			}
		}
		for k, v := range session.recoveryQueue {
			if v.sendTime.Add(time.Second * 8).Before(time.Now()) {
				session.packetToSend[k] = v
				delete(session.recoveryQueue, k)
			} else {
				break
			}
		}
		for k := range session.receivedWindow {
			if int64(k) < session.windowStart {
				delete(session.receivedWindow, k)
			} else {
				break
			}
		}
		session.processSendQueue()
	}
}

func (session *Session) processSendQueue() {
	if len(session.sendQueue.Packets) > 0 {
		session.sendQueue.SeqNumber = session.sendSeq
		session.sendSeq++
		session.sendDataPacket(session.sendQueue)
		pk := packet.NewPacket(session.sendQueue.Head)
		pk.Buffer = session.sendQueue.Buffer
		session.recoveryQueue[session.sendQueue.SeqNumber] = packetRecovery{
			sendTime:  time.Now(),
			seqNumber: session.sendQueue.SeqNumber,
			packet:    pk.GetBytes(),
		}
		session.sendQueue = packet.DataPacket{
			Buffer:              new(bytes.Buffer),
			SeqNumber:           0,
			Head:                0x84,
			EncapsulatedPackets: make([]packet.EncapsulatedPacket, 0),
			Packets:             make([][]byte, 0),
		}
	}
}

//Do not encode EncapsulatedPacket before sending. It will be done automatically.
func (session *Session) addToQueue(ep packet.EncapsulatedPacket, immediate bool) {
	if err := ep.Encapsulate(); err != nil {
		logger.FromError(err, 0)
		return
	}
	if immediate {
		dpk := packet.DataPacket{
			Buffer:    new(bytes.Buffer),
			Head:      0x80,
			SeqNumber: session.sendSeq,
		}
		dpk.Packets = [][]byte{ep.Bytes()}
		pk := dpk.Encode()
		session.SendStream <- pk.GetBytes()
		session.recoveryQueue[session.sendSeq] = packetRecovery{
			sendTime:  time.Now(),
			seqNumber: session.sendSeq,
			packet:    pk.GetBytes(),
		}
		session.sendSeq++
		return
	}
	length := session.sendQueue.TotalLen()
	if length+ep.TotalLen() > int(session.mtuSize) {
		session.processSendQueue()
	}
	session.sendQueue.Packets = append(session.sendQueue.Packets, ep.Bytes())
}

func (session *Session) addEncapsulatedToQueue(ep packet.EncapsulatedPacket, immediate bool) {
	if ep.Reliability != 0 && ep.Reliability != 1 && ep.Reliability != 5 {
		ep.MessageIndex = session.messageIndex
		session.messageIndex++
		if ep.Reliability == 3 {
			ep.OrderIndex = session.channelIndex[uint16(ep.OrderChannel)]
			session.channelIndex[uint16(ep.OrderChannel)]++
		}
	}
	if ep.TotalLen()+4 > int(session.mtuSize) {
		buffer := ep.Bytes()
		splitID := session.splitID
		session.splitID++
		start, end := uint64(0), session.mtuSize-34
		for cnt := 0; int(start) < len(buffer); cnt++ {
			if int(end) > len(buffer) {
				end = uint64(len(buffer))
			}
			split := buffer[start:end]
			pk := *new(packet.EncapsulatedPacket)
			pk.SplitID = splitID
			pk.HasSplit = true
			pk.SplitCount = uint32(len(buffer)/int(session.mtuSize-34)) + 1
			pk.Reliability = ep.Reliability
			pk.SplitIndex = uint32(cnt)
			pk.Buffer = bytes.NewBuffer(split)
			if cnt > 0 {
				pk.MessageIndex = session.messageIndex
				session.messageIndex++
			} else {
				pk.MessageIndex = session.messageIndex
			}
			if pk.Reliability == 3 {
				pk.OrderChannel = ep.OrderChannel
				pk.OrderIndex = ep.OrderIndex
			}
			session.addToQueue(ep, true)
			start += session.mtuSize - 34
			end += session.mtuSize - 34
		}
	} else {
		session.addToQueue(ep, immediate)
	}
}

func (session *Session) handlePacket(pk packet.Packet) {
	session.isActive = true
	session.lastUpdate = time.Now()
	if session.state == connected || session.state == connecting2 {
		if pk.Head >= 0x80 && pk.Head <= 0x8f { // Data packet
			var dpk packet.DataPacket
			var err error
			if dpk, err = packet.FromPacket(pk); err != nil {
				logger.FromError(err, 0)
				return
			}
			if _, ok := session.receivedWindow[dpk.SeqNumber]; int64(dpk.SeqNumber) < session.windowStart || int64(dpk.SeqNumber) > session.windowEnd || ok {
				return
			}
			diff := int64(dpk.SeqNumber - session.lastSeq)
			if _, ok := session.nackQueue[dpk.SeqNumber]; ok {
				delete(session.nackQueue, dpk.SeqNumber)
			}
			session.ackQueue[dpk.SeqNumber] = dpk.SeqNumber
			session.receivedWindow[dpk.SeqNumber] = dpk
			if diff != 1 {
				for i := session.lastSeq + 1; i < dpk.SeqNumber; i++ {
					if _, ok := session.receivedWindow[i]; !ok {
						session.nackQueue[i] = i
					}
				}
			}
			if diff >= 1 {
				session.lastSeq = dpk.SeqNumber
				session.windowStart += diff
				session.windowEnd += diff
			}
			for _, v := range dpk.EncapsulatedPackets {
				session.handleEncapsulatedPacket(v)
			}
		} else {
			if pk.Head == 0xc0 { // ACK
				ack := packet.AcknowledgePacket{Buffer: bytes.NewBuffer(pk.Bytes())}
				ack.Decode()
				for _, v := range ack.Packets {
					if _, ok := session.recoveryQueue[v]; ok {
						delete(session.recoveryQueue, v)
					}
				}
			} else if pk.Head == 0xa0 { // NACK
				nack := packet.AcknowledgePacket{Buffer: bytes.NewBuffer(pk.Bytes())}
				nack.Decode()
				for _, v := range nack.Packets {
					if rpk, ok := session.recoveryQueue[v]; ok {
						seq := session.sendSeq
						session.sendSeq++
						session.packetToSend[seq] = packetRecovery{packet: append([]byte{0x84, byte(seq >> 16), byte(seq >> 8), byte(seq)}, rpk.packet[4:]...)}
						delete(session.recoveryQueue, seq)
					}
				}
			}
		}
	} else if pk.Head < 0x80 { // Not data packet
		if pk.Head == 0x05 {
			pk.Next(16) //Magic
			if len(pk.Buffer.Bytes()) < 18 {
				logger.Error("Error while processing packet parse: buffer too short")
			} else if proto, err := pk.ReadByte(); err == nil && int(proto) != RaknetProtocol {
				logger.Error("Raknet protocol mismatch: " + strconv.Itoa(int(pk.Buffer.Bytes()[16])) + " != " + strconv.Itoa(RaknetProtocol))
			} else if err != nil {
				logger.FromError(err, 0)
			}
			mtusize := make([]byte, 2)
			if _, err := pk.Read(mtusize); err != nil {
				logger.FromError(err, 0)
			}
			session.mtuSize = uint64(math.Min(float64(binary.BigEndian.Uint16(mtusize)+18), 1464))
			pk := packet.NewPacket(0x06)
			pk.Buffer.Write([]byte(RaknetMagic))
			binary.Write(pk, binary.BigEndian, session.ServerID)
			pk.WriteByte(0)
			binary.Write(pk, binary.BigEndian, mtusize)
			session.SendStream <- pk.GetBytes()
			session.state = connecting1
			logger.Debug("set state to 1")
		} else if session.state == connecting1 && pk.Head == 0x07 {
			pk.Next(16) //Magic
			if _, err := packet.ReadAddress(pk.Buffer); err != nil {
				logger.FromError(err, 0)
				logger.Debug(hex.EncodeToString([]byte{pk.Head}), "\n"+hex.Dump(append([]byte{pk.Head}, pk.Buffer.Bytes()...)))
			}
			pk = packet.NewPacket(0x08)
			pk.Write([]byte(RaknetMagic))
			binary.Write(pk.Buffer, binary.BigEndian, session.ServerID)
			packet.PutAddress(session.Address, pk.Buffer, 4)
			binary.Write(pk.Buffer, binary.BigEndian, session.mtuSize)
			session.SendStream <- pk.GetBytes()
			session.state = connecting2
			logger.Debug("set state to 2")
		}
	}
}

func (session *Session) handleEncapsulatedPacket(ep packet.EncapsulatedPacket) {
	if ep.Reliability >= 2 && ep.Reliability != 5 {
		if ep.MessageIndex < session.reliableWindowStart || ep.MessageIndex > session.reliableWindowEnd {
			return
		}
		if (int64(ep.MessageIndex) - session.lastReliableIndex) == 1 {
			session.lastReliableIndex++
			session.reliableWindowStart++
			session.reliableWindowEnd++
			session.handleEncapsulatedPacketRoute(ep)
			if len(session.reliableWindow) > 0 {
				var keys []int
				for k := range session.reliableWindow {
					keys = append(keys, int(k))
				}
				sort.Ints(keys)
				for i := range keys {
					if (int64(i) - session.lastReliableIndex) != 1 {
						break
					}
					session.lastReliableIndex++
					session.reliableWindowStart++
					session.reliableWindowEnd++
					session.handleEncapsulatedPacketRoute(session.reliableWindow[uint32(i)])
					delete(session.reliableWindow, uint32(i))
				}
			}
		} else {
			session.reliableWindow[ep.MessageIndex] = ep
		}
	} else {
		session.handleEncapsulatedPacketRoute(ep)
	}
}

func (session *Session) handleEncapsulatedPacketRoute(ep packet.EncapsulatedPacket) {
	if ep.HasSplit {
		if session.state == connected {
			session.handleSplit(ep)
		}
		return
	}
	head := ep.Payload[0]
	if head < 0x80 { // Internal raknet signal
		//TODO
	} else if session.state == connected {
		//TODO: stream datapackets to player handler
	} else {
		logger.Debug("?! Received datapacket before connection: Dumping decapsulated payload\n" + hex.Dump(ep.Payload))
	}
}

func (session *Session) handleSplit(ep packet.EncapsulatedPacket) {
	if ep.SplitCount >= maxSplitSize || ep.SplitIndex >= maxSplitSize {
		return
	}
	if _, ok := session.splitPackets[ep.SplitID]; !ok {
		if len(session.splitPackets) >= maxSplitCount {
			return
		}
		session.splitPackets[ep.SplitID] = map[uint32][]byte{
			ep.SplitIndex: ep.Bytes(),
		}
	} else {
		session.splitPackets[ep.SplitID][ep.SplitIndex] = ep.Bytes()
	}
	if uint32(len(session.splitPackets[ep.SplitID])) == ep.SplitCount {
		pk := *new(packet.EncapsulatedPacket)
		pk.Buffer = new(bytes.Buffer)
		for i := uint32(0); i < ep.SplitCount; i++ {
			pk.Write(session.splitPackets[ep.SplitID][i])
		}
		delete(session.splitPackets, ep.SplitID)
		session.handleEncapsulatedPacketRoute(pk)
	}
}

func (session *Session) sendDataPacket(dpk packet.DataPacket) {
	session.SendStream <- append([]byte{dpk.Head}, dpk.Bytes()...)
}

func (session *Session) disconnect(reason string) {
	logger.Verbose("Closing session", session.Address, "(Reason: "+reason+")")
	session.addEncapsulatedToQueue(packet.EncapsulatedPacket{Buffer: bytes.NewBuffer([]byte("\x00\x00\x08\x15")), Locked: true}, true)
	close(session.SendStream)
	close(session.RecvStream)
	runtime.Goexit()
}
