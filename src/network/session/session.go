package session

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
	MtuSize         uint16
	ConnectionState byte
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
			if pk.Head >= 0x80 && pk.Head < 0x90 {
				session.handleDataPacket(pk)
			}
			switch {
			case session.ConnectionState == unconnected && pk.Head == 0x05:
				if len(pk.Buffer.Bytes()) < 18 {
					fmt.Println("Packet error from", pk.Address, "(buffer too short)")
					continue
				} else if int(pk.Buffer.Bytes()[16]) != protocol.RaknetProtocol {
					fmt.Println("Raknet protocol mismatch:", pk.Buffer.Bytes()[16], "!=", protocol.RaknetProtocol)
					continue
				}
				session.ConnectionState = connecting1
				fmt.Println(pk.Address, "set state to 1")
				mtusize := make([]byte, 2)
				if _, err := pk.Read(mtusize); err != nil {
					fmt.Println("Unexpected packet error from", pk.Address, ":", err)
					continue
				}
				session.MtuSize = binary.BigEndian.Uint16(mtusize)
				pk := packet.NewPacket(0x06)
				pk.Buffer.Write(protocol.RaknetMagic)
				binary.Write(pk, binary.BigEndian, session.ServerID)
				pk.WriteByte(0)
				binary.Write(pk, binary.BigEndian, mtusize)
				session.SendStream <- pk
			case session.ConnectionState == connecting1 && pk.Head == 0x07:

			case session.ConnectionState == connecting2:
			case session.ConnectionState == connected:
			}
			fmt.Print("Unexpected packet header: 0x", hex.EncodeToString([]byte{pk.Head}), "\n"+hex.Dump(append([]byte{pk.Head}, pk.Buffer.Bytes()...)))
		}
	}
}

func (session *Session) asyncTiming() {

}

func (session *Session) handleDataPacket(pk packet.Packet) {

}

//Close Closes session channels for stopping goroutines
func (session *Session) Close() {
	//TODO: Implement close packet
	close(session.RecvStream)
	close(session.SendStream)
}
