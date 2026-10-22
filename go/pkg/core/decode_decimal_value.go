//ff:func feature=core type=decoder control=selection
//ff:what DecodeDecimalValue — the exponent and coefficient of a SIGNED-/UNSIGNED-DECIMAL Value record
package core

import (
	"errors"
	"fmt"
)

// DecodeDecimalValue returns the exponent E and coefficient C of r, a
// SIGNED-DECIMAL or UNSIGNED-DECIMAL Value record (value C×10^E). The payload is
// a 2-word signed exponent plus at least one minimal coefficient word; an
// unsigned coefficient must fit int64.
func DecodeDecimalValue(r Record) (int32, int64, error) {
	if r.Kind != KindValue || len(r.Payload) < 3 {
		return 0, 0, errors.New("geul/core: decimal Value needs exponent and coefficient words")
	}
	e := int32(uint32(r.Payload[0])<<16 | uint32(r.Payload[1]))
	switch r.Codec {
	case CodecSignedDecimal:
		c, err := wordsToInt(r.Payload[2:])
		return e, c, err
	case CodecUnsignedDecimal:
		u, err := wordsToUint(r.Payload[2:])
		if err == nil && u > 1<<63-1 {
			err = errors.New("geul/core: UNSIGNED-DECIMAL coefficient exceeds int64")
		}
		return e, int64(u), err
	default:
		return 0, 0, fmt.Errorf("geul/core: expected a decimal Value, got codec 0x%02X", r.Codec)
	}
}
