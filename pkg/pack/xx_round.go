//ff:func feature=pack type=helper control=sequence
//ff:what xxRound — one XXH64 accumulator round
package pack

import "math/bits"

// xxRound mixes one 64-bit lane into acc.
func xxRound(acc, in uint64) uint64 {
	acc += in * xxP2
	acc = bits.RotateLeft64(acc, 31)
	return acc * xxP1
}
