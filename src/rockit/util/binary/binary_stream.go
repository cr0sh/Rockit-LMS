package binary

import "encoding/binary"

//Stream reads/writes packet buffer, providing binary functions
type Stream struct {
	Buffer []byte
	Offset uint
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//Read reads stream with given length
func (s *Stream) Read(n uint) (buf []byte) {
	buf = s.Buffer[s.Offset:min(int(s.Offset+n), len(s.Buffer))]
	s.Offset = uint(min(int(s.Offset+n), len(s.Buffer)))
	return
}

//Write writes byte array to stream
func (s *Stream) Write(buf []byte) {
	s.Buffer = append(s.Buffer, buf...)
}

//ReadShort reads unsigned 16-bit integer from packet
func (s *Stream) ReadShort() uint16 {
	return binary.BigEndian.Uint16(s.Read(2))
}

//ReadInt reads unsigned 32-bit integer from packet
func (s *Stream) ReadInt() uint32 {
	return binary.BigEndian.Uint32(s.Read(4))
}

//ReadLong reads unsigned 64-bit integer from packet
func (s *Stream) ReadLong() uint64 {
	return binary.BigEndian.Uint64(s.Read(8))
}

//ReadString reads string from packet
func (s *Stream) ReadString() string {
	len := s.ReadShort()
	return string(s.Read(uint(len)))
}

//WriteShort wirtes unsigned 16-bit integer to packet
func (s *Stream) WriteShort(n uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, n)
	s.Buffer = append(s.Buffer, buf...)
}

//WriteInt wirtes unsigned 32-bit integer to packet
func (s *Stream) WriteInt(n uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n)
	s.Buffer = append(s.Buffer, buf...)
}

//WriteLong wirtes unsigned 64bit integer to packet
func (s *Stream) WriteLong(n uint64) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	s.Buffer = append(s.Buffer, buf...)
}

//WriteString writes string to packet
func (s *Stream) WriteString(str string) {
	s.WriteShort(uint16(len(str)))
	s.Write([]byte(str))
}
