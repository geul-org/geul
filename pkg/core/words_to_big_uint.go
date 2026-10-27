//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what wordsToBigUint — minimal big-endian unsigned WORDs of any length as a non-negative integer
package core

import (
	"errors"
	"math/big"
)

// wordsToBigUint is wordsToUint for any length: ws read as a big-endian
// unsigned integer. It needs at least one WORD and rejects a leading 0x0000
// WORD on a longer run (minimal form).
func wordsToBigUint(ws []Word) (*big.Int, error) {
	if len(ws) == 0 {
		return nil, errors.New("geul/core: unsigned integer needs at least one word")
	}
	if len(ws) > 1 && ws[0] == 0 {
		return nil, errors.New("geul/core: non-minimal unsigned integer (leading 0x0000)")
	}
	n := new(big.Int)
	for _, w := range ws {
		n.Lsh(n, 16).Or(n, big.NewInt(int64(w)))
	}
	return n, nil
}
