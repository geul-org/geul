//ff:func feature=pack type=decoder control=sequence
//ff:what dictID — the Dictionary_ID inside a formatted zstd dictionary (magic 0xEC30A437); false for raw content
package pack

import "encoding/binary"

// dictID reads a formatted dictionary's header (RFC 8878 5). Bytes without
// the magic are a raw-content dictionary, which carries no ID.
func dictID(b []byte) (uint32, bool) {
	if len(b) < 8 || binary.LittleEndian.Uint32(b) != dictMagic {
		return 0, false
	}
	return binary.LittleEndian.Uint32(b[4:]), true
}
