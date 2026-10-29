//ff:func feature=pack type=writer control=sequence
//ff:what appendHEntry — one H run entry as hEntrySize bytes: key, object, offset, length, encoding
package pack

import "encoding/binary"

// appendHEntry appends e in run format (little-endian numbers).
func appendHEntry(b []byte, e hEntry) []byte {
	b = append(b, e.key[:]...)
	b = binary.LittleEndian.AppendUint32(b, e.loc.Obj)
	b = binary.LittleEndian.AppendUint64(b, e.loc.Off)
	b = binary.LittleEndian.AppendUint64(b, e.loc.Len)
	return append(b, e.loc.Enc)
}
