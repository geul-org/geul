//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what wordsToBigInt — minimal big-endian two's complement WORDs of any length as a signed integer
package core

import (
	"errors"
	"math/big"
)

// wordsToBigInt is wordsToInt for any length: ws read as a big-endian two's
// complement integer. It needs at least one WORD and rejects a removable
// leading sign WORD (minimal form).
func wordsToBigInt(ws []Word) (*big.Int, error) {
	if len(ws) == 0 {
		return nil, errors.New("geul/core: signed integer needs at least one word")
	}
	if len(ws) > 1 && (ws[0] == 0 && ws[1]&0x8000 == 0 || ws[0] == 0xFFFF && ws[1]&0x8000 != 0) {
		return nil, errors.New("geul/core: non-minimal signed integer (removable sign word)")
	}
	n := big.NewInt(int64(int16(ws[0])))
	for _, w := range ws[1:] {
		n.Lsh(n, 16).Or(n, big.NewInt(int64(w)))
	}
	return n, nil
}
