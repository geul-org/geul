//ff:func feature=core type=decoder control=sequence
//ff:what DecodeSignedDecimalBig — the exponent and coefficient of a SIGNED-DECIMAL Value of any coefficient size
package core

import (
	"errors"
	"math/big"
)

// DecodeSignedDecimalBig returns E and C of r, a SIGNED-DECIMAL Value (C×10^E)
// whose coefficient may exceed int64. The coefficient must be minimal.
func DecodeSignedDecimalBig(r Record) (int32, *big.Int, error) {
	if r.Kind != KindValue || r.Codec != CodecSignedDecimal || len(r.Payload) < 3 {
		return 0, nil, errors.New("geul/core: expected a SIGNED-DECIMAL Value with exponent and coefficient words")
	}
	c, err := wordsToBigInt(r.Payload[2:])
	return int32(uint32(r.Payload[0])<<16 | uint32(r.Payload[1])), c, err
}
