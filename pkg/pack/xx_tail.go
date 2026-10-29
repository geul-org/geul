//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what xxTail — XXH64 tail processing (8-, 4- and 1-byte steps) and the final avalanche
package pack

import (
	"encoding/binary"
	"math/bits"
)

// xxTail finishes h over the remaining bytes (fewer than 32).
func xxTail(h uint64, b []byte) uint64 {
	for len(b) >= 8 {
		h ^= xxRound(0, binary.LittleEndian.Uint64(b))
		h = bits.RotateLeft64(h, 27)*xxP1 + xxP4
		b = b[8:]
	}
	if len(b) >= 4 {
		h ^= uint64(binary.LittleEndian.Uint32(b)) * xxP1
		h = bits.RotateLeft64(h, 23)*xxP2 + xxP3
		b = b[4:]
	}
	for _, c := range b {
		h ^= uint64(c) * xxP5
		h = bits.RotateLeft64(h, 11) * xxP1
	}
	h ^= h >> 33
	h *= xxP2
	h ^= h >> 29
	h *= xxP3
	return h ^ h>>32
}
