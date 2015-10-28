package session

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"network/packet"
	"network/protocol"
	"strconv"
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
	sendSeqNum      int32
}

const (
	unconnected byte = 0
	connecting1 byte = 1
	connecting2 byte = 2
	connected   byte = 3
)

//Handle Handles packets from RecvStream channel
func (session *Session) Handle() {
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
	if pk.Head >= 0x80 && pk.Head < 0x90 {
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

func (session *Session) asyncProcess() { //TODO

}

func (session *Session) handleDataPacket(dp packet.DataPacket) bool {
	logging.Debug("Seqnumber " + strconv.Itoa(int(dp.SeqNumber)))
	if dp.SeqNumber > 10000 {
		logging.Debug("\n" + hex.Dump(append([]byte{dp.Head}, dp.Bytes()...)))
	}
	for _, pk := range dp.Packets {
		session.handleEncapsulatedPacket(pk)
	}
	return true
}

func (session *Session) handleEncapsulatedPacket(pk packet.Packet) {
	logging.Debug("DataPacket head 0x" + hex.EncodeToString([]byte{pk.Head}))
	switch pk.Head {
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
		ep.HasSplit = false
		ep.Encapsulate(pk)
		session.sendDataPacket(ep)
	}
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
func (session *Session) Close() {
	//TODO: Implement close packet
	close(session.RecvStream)
	close(session.SendStream)
}
