//Package socket provides asynchronous UDP socket and sends packets to each sessions
package socket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"rockit/network/packet"
	"rockit/network/protocol"
	"rockit/network/session"
	"rockit/util/logging"
	"strconv"
	"strings"
)

//Socket struct.
type Socket struct {
	ServerConn *net.UDPConn
	Input      chan packet.Packet
	Sessions   map[string]*session.Session
}

//ServerID variable
var ServerID uint64

//Open opens socket with given port
func (s *Socket) Open(port int16) (err error) {
	logging.Verbose("Opening socket on 0.0.0.0:", port)
	s.Input = make(chan packet.Packet, 1024)
	ServerAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+strconv.Itoa(int(port)))
	if err != nil {
		return err
	}
	s.ServerConn = new(net.UDPConn)
	s.ServerConn, err = net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return err
	}
	s.Sessions = make(map[string]*session.Session)
	return nil
}

//ProcessRecv receives packet asynchronously and sends packets to sessions.
func (s *Socket) ProcessRecv() {
	buffer := make([]byte, 1024*1024*8) //Buffer size 8MB
	go s.sendPackets()
	for {
		if n, addr, err := s.ServerConn.ReadFromUDP(buffer); err == nil {
			pk := &packet.Packet{Buffer: bytes.NewBuffer(buffer[1:n]), Head: buffer[0], Address: *addr}
			if pk.Head == 0x01 {
				var PingID uint64
				if err := binary.Read(pk.Buffer, binary.BigEndian, &PingID); err == nil {
					pk = new(packet.Packet)
					pk.Address = *addr
					pk.Buffer = new(bytes.Buffer)
					pk.Head = 0x1c
					binary.Write(pk.Buffer, binary.BigEndian, PingID)
					binary.Write(pk.Buffer, binary.BigEndian, ServerID)
					binary.Write(pk.Buffer, binary.BigEndian, protocol.RaknetMagic)
					pk.PutStr("MCPE;Rockit - using dev build now;34;" + protocol.MinecraftVersion + ";0;20")
					s.sendPacket(*pk)
				} else {
					fmt.Print("Error while decoding packet:", err)
				}
				continue
			}
			s.sendToSession(pk)
		} else {
			fmt.Println("Error:", err)
		}
	}

}

func (s *Socket) sendPacket(pk packet.Packet) {
	s.ServerConn.WriteToUDP(append([]byte{pk.Head}, pk.Buffer.Bytes()...), &pk.Address)
}

//ProcessSend gets a packet from Socket.Input channel and sends it
func (s *Socket) ProcessSend() {
	for snd := range s.Input {
		s.sendPacket(snd)
	}
}

func (s *Socket) sendToSession(pk *packet.Packet) {
	s.getSession(pk.Address).RecvStream <- *pk
}

func (s *Socket) getSession(address net.UDPAddr) *session.Session {
	addr := address.IP.String() + ":" + strconv.Itoa(address.Port)
	if sess, ok := s.Sessions[addr]; ok {
		return sess
	}
	fmt.Println("New session:", addr)
	sess := &session.Session{
		Address:    address,
		RecvStream: make(chan packet.Packet, 1024),
		SendStream: make(chan packet.Packet, 1024),
		ServerID:   ServerID,
	}
	s.Sessions[addr] = sess
	go sess.Handle()
	return sess

}

func (s *Socket) sendPackets() {
	for {
		for addr, sess := range s.Sessions {
			select {
			case pk, ok := <-sess.SendStream:
				if !ok {
					delete(s.Sessions, addr)
					continue
				}
				address := net.ParseIP(strings.Split(addr, ":")[0])
				port, _ := strconv.Atoi(strings.Split(addr, ":")[1])
				pk.Address = net.UDPAddr{IP: address, Port: port, Zone: ""}
				s.Input <- pk
			default:
				continue
			}
		}
	}
}
