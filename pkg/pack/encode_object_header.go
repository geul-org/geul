//ff:func feature=pack type=encoder control=sequence
//ff:what encodeObjectHeader — the 64-byte object head
package pack

import "encoding/binary"

// encodeObjectHeader writes h as 64 bytes.
func encodeObjectHeader(h objectHeader) []byte {
	out := []byte(objectMagic)
	out = binary.LittleEndian.AppendUint16(out, formatVersion)
	out = binary.LittleEndian.AppendUint16(out, uint16(h.kind))
	out = binary.LittleEndian.AppendUint32(out, 0)
	out = binary.LittleEndian.AppendUint64(out, h.size)
	out = binary.LittleEndian.AppendUint64(out, h.invOff)
	out = binary.LittleEndian.AppendUint64(out, h.invLen)
	out = binary.LittleEndian.AppendUint64(out, h.count)
	return append(out, make([]byte, 16)...)
}
