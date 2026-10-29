//ff:func feature=pack type=encoder control=sequence
//ff:what pKey — a P key (Issuer number, document ID, slot) as 16 big-endian bytes so byte order is key order
package pack

import "encoding/binary"

// pKey packs (issuer number, document ID, slot) into a 16-byte string whose
// byte order equals the numeric key order.
func pKey(i uint32, d uint64, s uint32) string {
	b := make([]byte, 0, 16)
	b = binary.BigEndian.AppendUint32(b, i)
	b = binary.BigEndian.AppendUint64(b, d)
	b = binary.BigEndian.AppendUint32(b, s)
	return string(b)
}
