//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what EncodeID — write an unsigned id/offset as a fixed-width big-endian word field
package core

import "fmt"

// EncodeID writes v as a big-endian field of exactly width words (1, 2, 4 or 8)
// — the declared-width form of Edge Ids, references and SPAN offsets. A value
// that does not fit the width is an error, never a silent truncation.
func EncodeID(v uint64, width int) ([]Word, error) {
	if _, err := WidthCode(width); err != nil {
		return nil, err
	}
	if width < 4 && v>>(16*uint(width)) != 0 {
		return nil, fmt.Errorf("geul/core: value %d does not fit %d words", v, width)
	}
	out := make([]Word, width)
	for i := width - 1; i >= 0; i-- {
		out[i] = Word(v)
		v >>= 16
	}
	return out, nil
}
