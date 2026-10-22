//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what wordsToUint — read a minimal big-endian unsigned word run, rejecting padding and uint64 overflow
package core

import "errors"

// wordsToUint reads ws as a minimal big-endian unsigned integer: it must be at
// least one word, carry no leading 0x0000 word when longer than one word, and
// fit in uint64 (at most four words once minimal).
func wordsToUint(ws []Word) (uint64, error) {
	if len(ws) == 0 {
		return 0, errors.New("geul/core: unsigned integer needs at least one word")
	}
	if len(ws) > 1 && ws[0] == 0 {
		return 0, errors.New("geul/core: non-minimal unsigned integer (leading 0x0000)")
	}
	if len(ws) > 4 {
		return 0, errors.New("geul/core: unsigned integer exceeds 64 bits")
	}
	var n uint64
	for _, w := range ws {
		n = n<<16 | uint64(w)
	}
	return n, nil
}
