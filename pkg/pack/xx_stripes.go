//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what xxStripes — XXH64 over the 32-byte stripes of an input of 32 bytes or more; returns the rest
package pack

import (
	"encoding/binary"
	"math/bits"
)

// xxStripes consumes whole 32-byte stripes with four accumulators and returns
// the merged hash and the unconsumed tail.
func xxStripes(b []byte) (uint64, []byte) {
	p1, p2 := xxP1, xxP2
	v1, v2, v3, v4 := p1+p2, p2, uint64(0), -p1
	for len(b) >= 32 {
		v1 = xxRound(v1, binary.LittleEndian.Uint64(b))
		v2 = xxRound(v2, binary.LittleEndian.Uint64(b[8:]))
		v3 = xxRound(v3, binary.LittleEndian.Uint64(b[16:]))
		v4 = xxRound(v4, binary.LittleEndian.Uint64(b[24:]))
		b = b[32:]
	}
	h := bits.RotateLeft64(v1, 1) + bits.RotateLeft64(v2, 7) + bits.RotateLeft64(v3, 12) + bits.RotateLeft64(v4, 18)
	h = xxMerge(xxMerge(xxMerge(xxMerge(h, v1), v2), v3), v4)
	return h, b
}
