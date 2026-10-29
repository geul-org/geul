//ff:func feature=pack type=decoder control=sequence
//ff:what splitPKey — the Issuer number, document ID and slot of a 16-byte P key
package pack

import "encoding/binary"

// splitPKey is the inverse of pKey.
func splitPKey(k string) (uint32, uint64, uint32) {
	b := []byte(k)
	return binary.BigEndian.Uint32(b), binary.BigEndian.Uint64(b[4:]), binary.BigEndian.Uint32(b[12:])
}
