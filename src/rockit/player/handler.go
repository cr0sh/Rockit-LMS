package player

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"
	"rockit/network/packet/mcpe"
	"rockit/util"
)

//Handler handles packets from player and controls player entity.
type Handler struct {
	Address  net.UDPAddr
	entity   Player
	username string
}

//HandlePacket handles MCPE DataPacket from player
func (handler *Handler) HandlePacket(pk []byte) {
	util.Debug("Handling MCPE Packet: head 0x" + hex.EncodeToString([]byte{pk[0]}))
	var ppk mcpe.Packet
	var err error
	ppk, err = mcpe.GetPacket(pk[0])
	if err != nil {
		util.FromError(err, 1)
		return
	}
	if fields, err := ppk.Decode(bytes.NewBuffer(pk[1:])); err == nil { //DO NOT PASS ENTIRE BUFFER. Skip head.
		switch pk[0] {
		case mcpe.BatchPacketHead:
			input := bytes.NewBuffer(fields["payload"].([]byte))
			r, err := zlib.NewReader(input)
			if err != nil {
				util.FromError(err, 0)
				return
			}
			output := new(bytes.Buffer)
			io.Copy(output, r)
			r.Close()
			buf := bytes.NewBuffer(output.Bytes())
			maxlen := uint32(buf.Len())
			offset := uint32(0)
			for offset < maxlen {
				tmp := binary.BigEndian.Uint32(buf.Next(4))
				offset += 4
				dpc := buf.Next(int(tmp))
				offset += tmp
				if dpc[0] == mcpe.BatchPacketHead {
					util.Error("Invalid BatchPacket inside BatchPacket")
					return
				}
				handler.HandlePacket(dpc)
			}
		}
	} else {
		util.FromError(err, 0)
	}
}
