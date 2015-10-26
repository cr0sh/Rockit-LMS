package session

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"network/packet"
	"network/protocol"
)

//Session struct contains packet channels for packet I/O, connection state, and source address.
type Session struct {
	Address         net.UDPAddr
	RecvStream      chan packet.Packet
	SendStream      chan packet.Packet
	ServerID        uint64
	mtuSize         uint16
	connectionState byte
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
		session.handleDataPacket(pk)
	}
	switch {
	case pk.Head == 0x05:
		pk.Next(16) //Magic
		if len(pk.Buffer.Bytes()) < 18 {
			fmt.Println("Packet error from", pk.Address, "(buffer too short)")
			return true
		} else if proto, err := pk.ReadByte(); err == nil && int(proto) != protocol.RaknetProtocol {
			fmt.Println("Raknet protocol mismatch:", pk.Buffer.Bytes()[16], "!=", protocol.RaknetProtocol)
			return true
		} else if err != nil {
			fmt.Println("Unexpected packet error from ", pk.Address, ":", err)
			return true
		}
		mtusize := make([]byte, 2)
		if _, err := pk.Read(mtusize); err != nil {
			fmt.Println("Unexpected packet error from", session.Address, ":", err)
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
			fmt.Println("Unexpected packet error from", session.Address, ":", err)
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
		fmt.Println("set state to 2")
		return true
	case session.connectionState == connecting2 || session.connectionState == connected:
		return session.handleDataPacket(pk)
	}
	return false
}

func (session *Session) asyncProcess() { //TODO

}

func (session *Session) handleDataPacket(pk packet.Packet) bool {

	return false
}

//Close Closes session channels for stopping goroutines
func (session *Session) Close() {
	//TODO: Implement close packet
	close(session.RecvStream)
	close(session.SendStream)
}
