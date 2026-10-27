//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what DecodeID — read a fixed-width big-endian id/offset field into a uint64
package core

import (
	"errors"
	"fmt"
)

// DecodeID reads the first width words (1, 2, 4 or 8) of ws as a big-endian
// unsigned value. An 8-word field whose value exceeds 64 bits is valid GEUL but
// beyond this implementation, so it is reported rather than truncated.
func DecodeID(ws []Word, width int) (uint64, error) {
	if _, err := WidthCode(width); err != nil {
		return 0, err
	}
	if len(ws) < width {
		return 0, fmt.Errorf("geul/core: %d-word field truncated", width)
	}
	var v uint64
	for i, w := range ws[:width] {
		if width-i > 4 && w != 0 {
			return 0, errors.New("geul/core: field value exceeds 64 bits")
		}
		v = v<<16 | uint64(w)
	}
	return v, nil
}
