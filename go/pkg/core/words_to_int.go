//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what wordsToInt — read a minimal two's-complement word run, rejecting removable sign words and int64 overflow
package core

import "errors"

// wordsToInt reads ws as a minimal big-endian two's-complement integer. A
// leading 0x0000 (or 0xFFFF) word followed by a word with the same sign bit is
// removable and therefore rejected; more than four words cannot fit int64.
func wordsToInt(ws []Word) (int64, error) {
	if len(ws) == 0 {
		return 0, errors.New("geul/core: signed integer needs at least one word")
	}
	if len(ws) > 1 && (ws[0] == 0 && ws[1]&0x8000 == 0 || ws[0] == 0xFFFF && ws[1]&0x8000 != 0) {
		return 0, errors.New("geul/core: non-minimal signed integer (removable sign word)")
	}
	if len(ws) > 4 {
		return 0, errors.New("geul/core: signed integer exceeds 64 bits")
	}
	n := int64(int16(ws[0]))
	for _, w := range ws[1:] {
		n = n<<16 | int64(w)
	}
	return n, nil
}
