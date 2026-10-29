//ff:func feature=pack type=writer control=sequence
//ff:what parseHEntry — one H run entry from its hEntrySize bytes
package pack

import "encoding/binary"

// parseHEntry is appendHEntry's inverse; b holds exactly hEntrySize bytes.
func parseHEntry(b []byte) hEntry {
	var e hEntry
	copy(e.key[:], b[:32])
	e.loc.Obj = binary.LittleEndian.Uint32(b[32:])
	e.loc.Off = binary.LittleEndian.Uint64(b[36:])
	e.loc.Len = binary.LittleEndian.Uint64(b[44:])
	e.loc.Enc = b[52]
	return e
}
