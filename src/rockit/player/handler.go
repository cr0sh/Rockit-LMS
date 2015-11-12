package player

import (
	"bytes"
	"compress/zlib"
	binutil "encoding/binary"
	"encoding/hex"
	"io"
	"net"
	"rockit/network/packet"
	"rockit/network/packet/mcpe"
	"rockit/util/binary"
	"rockit/util/logger"
)

//Handler handles packets from player and controls player entity.
type Handler struct {
	Address    net.UDPAddr
	SendStream chan packet.Packet
	entity     Player
	username   string
}

//HandlePacket handles MCPE DataPacket from player
func (handler *Handler) HandlePacket(pk []byte) {
	logger.Debug("Handling MCPE Packet: head 0x" + hex.EncodeToString([]byte{pk[0]}))
	var ppk mcpe.Packet
	var err error
	ppk, err = mcpe.GetPacket(pk[0])
	if err != nil {
		logger.FromError(err, 1)
		return
	}
	if fields, err := ppk.Decode(binary.Stream{pk[1:], 0}); err == nil { //DO NOT PASS ENTIRE BUFFER. Skip head.
		switch pk[0] {
		case mcpe.BatchPacketHead:
			input := bytes.NewBuffer(fields["payload"].([]byte))
			r, err := zlib.NewReader(input)
			if err != nil {
				logger.FromError(err, 0)
				return
			}
			output := new(bytes.Buffer)
			io.Copy(output, r)
			r.Close()
			buf := bytes.NewBuffer(output.Bytes())
			maxlen := uint32(buf.Len())
			offset := uint32(0)
			for offset < maxlen {
				tmp := binutil.BigEndian.Uint32(buf.Next(4))
				offset += 4
				dpc := buf.Next(int(tmp))
				offset += tmp
				if dpc[0] == mcpe.BatchPacketHead {
					logger.Error("Invalid BatchPacket inside BatchPacket")
					return
				}
				handler.HandlePacket(dpc)
			}
		case mcpe.LoginPacketHead:
			handler.username = fields["username"].(string)
			handler.sendPacket(mcpe.PlayStatusPacketHead, map[string]interface{}{"status": uint32(mcpe.Success)})
			return
		}
	} else {
		logger.FromError(err, 0)
	}
}

func (handler *Handler) sendPacket(head byte, fields mcpe.Field) {
	if pk, err := mcpe.GetPacket(head); err == nil {
		if buf, err := pk.Encode(fields); err == nil {
			handler.SendStream <- packet.Packet{Buffer: bytes.NewBuffer(buf), Head: head, Address: handler.Address}
		} else {
			logger.FromError(err, 1)
		}
	}
}
