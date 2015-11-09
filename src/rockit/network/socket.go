package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"rockit/network/packet"
	"rockit/util"
	"strconv"
	"strings"
)

//Socket struct.
type Socket struct {
	ServerConn *net.UDPConn
	Input      chan packet.Packet
	Sessions   map[string]*Session
	lastSID    uint
}

//ServerID variable
var ServerID uint64

//Open opens socket with given port
func (s *Socket) Open(port int16) (err error) {
	util.Verbose("Opening socket on 0.0.0.0:", port)
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
	s.Sessions = make(map[string]*Session)
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
					pk.Buffer = bytes.NewBuffer([]byte{})
					pk.Head = 0x1c
					if err := binary.Write(pk.Buffer, binary.BigEndian, PingID); err != nil {
						util.FromError(err, 0)
						continue
					}
					if err := binary.Write(pk.Buffer, binary.BigEndian, ServerID); err != nil {
						util.FromError(err, 0)
						continue
					}
					if err := binary.Write(pk.Buffer, binary.BigEndian, []byte(RaknetMagic)); err != nil {
						util.FromError(err, 0)
						continue
					}
					pk.PutStr("MCPE;Rockit - using dev build now;34;0.12.1;0;20")
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
	s.ServerConn.WriteToUDP(pk.GetBytes(), &pk.Address)
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

func (s *Socket) getSession(address net.UDPAddr) *Session {
	addr := address.IP.String() + ":" + strconv.Itoa(address.Port)
	if sess, ok := s.Sessions[addr]; ok {
		return sess
	}
	fmt.Println("New session:", addr)
	sess := &Session{
		Address:    address,
		RecvStream: make(chan packet.Packet, 1024),
		SendStream: make(chan packet.Packet, 1024),
		ServerID:   ServerID,
		SessionID:  s.lastSID,
	}
	s.lastSID++
	s.Sessions[addr] = sess
	go sess.HandleSession()
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
