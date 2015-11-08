package packet

import (
	"bytes"
	"encoding/binary"
	"sort"
)

type seqList []uint32

func (s seqList) Len() int {
	return len(s)
}

func (s seqList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s seqList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//AcknowledgePacket is a helper struct for encoding/decoding (N)ACK packets.
type AcknowledgePacket struct {
	*bytes.Buffer
	Packets seqList
}

//Encode encodes AcknowledgePacket.
func (a *AcknowledgePacket) Encode() {
	sort.Sort(a.Packets)
	payload := new(bytes.Buffer)
	count := a.Packets.Len()
	var records uint16
	if count > 0 {
		pointer, start, last := 1, a.Packets[0], a.Packets[0]
		for pointer < count {
			current := a.Packets[pointer]
			pointer++
			diff := current - last
			if diff == 1 {
				last = current
			} else if diff > 1 {
				if start == last {
					payload.WriteByte(0x01)
					PutLTriad(start, payload)
					last = current
					start = last
				} else {
					payload.WriteByte(0x00)
					PutLTriad(start, payload)
					PutLTriad(start, payload)
					last = current
					start = last
				}
				records++
			}
		}
		if start == last {
			payload.WriteByte(0x01)
			PutLTriad(start, payload)
		} else {
			payload.WriteByte(0x00)
			PutLTriad(start, payload)
			PutLTriad(start, payload)
		}
		records++
	}
	a.Buffer = new(bytes.Buffer)
	binary.Write(a.Buffer, binary.BigEndian, records)
	a.Write(payload.Bytes())
}

//Decode decodes AcknowledgePacket.
func (a *AcknowledgePacket) Decode() (err error) {
	var packetcount uint16
	if err = binary.Read(a.Buffer, binary.BigEndian, &packetcount); err != nil {
		return
	}
	cnt := 0
	a.Packets = make([]uint32, 0)
	var i uint16
	for ; i < packetcount && cnt < 4096; i++ {
		var flag byte
		if flag, err = a.ReadByte(); err != nil {
			return
		} else if flag == 0 {
			var start, end uint32
			if start, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			if end, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			if (end - start) > 512 {
				end = start + 512
			}
			for c := start; c <= end; c++ {
				a.Packets = append(a.Packets, c)
				cnt++
			}
		} else {
			var c uint32
			if c, err = ReadLTriad(a.Buffer); err != nil {
				return
			}
			a.Packets = append(a.Packets, c)
			cnt++
		}
	}
	return
}
