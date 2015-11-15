package network

import (
	"net"
	"rockit/network/packet"
	"time"
)

type packetRecovery struct {
	sendTime  time.Time
	seqNumber uint32
	packet    packet.Packet
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
	Address        net.UDPAddr
	ServerID       uint64
	SessionID      uint
	RecvStream     chan packet.Packet
	SendStream     chan packet.Packet
	packetToSend   map[uint32]packetRecovery
	recoveryQueue  map[uint32]packetRecovery
	ackQueue       []uint32
	nackQueue      []uint32
	receivedWindow map[uint32]packet.DataPacket
	windowStart    uint32
	windowEnd      uint32
}

//HandleSession is a core loop for session processing
func (session *Session) HandleSession() {
	session.packetToSend = make(map[uint32]packetRecovery)
	session.recoveryQueue = make(map[uint32]packetRecovery)
	session.ackQueue = make([]uint32, 0)
	session.nackQueue = make([]uint32, 0)
	session.receivedWindow = make(map[uint32]packet.DataPacket)
	update := time.NewTicker(time.Millisecond * 25)
	go session.updateJob(update.C)
	defer update.Stop()
}

func (session *Session) updateJob(tick <-chan time.Time) {
	for range tick {
		if len(session.ackQueue) > 0 {
			ack := new(packet.AcknowledgePacket)
			ack.Packets = session.ackQueue
			ack.Encode()
			pk := packet.NewPacket(0xc0)
			*pk.Buffer = *ack.Buffer
			session.SendStream <- pk
			session.ackQueue = make([]uint32, 0)
		}

		if len(session.nackQueue) > 0 {
			nack := new(packet.AcknowledgePacket)
			nack.Packets = session.ackQueue
			nack.Encode()
			pk := packet.NewPacket(0xa0)
			*pk.Buffer = *nack.Buffer
			session.SendStream <- pk
			session.nackQueue = make([]uint32, 0)
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
			if k < session.windowStart {
				delete(session.receivedWindow, k)
			} else {
				break
			}
		}
		session.sendQueue()
	}
}

func (session *Session) sendQueue() {

}
