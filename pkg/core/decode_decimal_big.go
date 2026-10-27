//ff:func feature=core type=decoder control=selection
//ff:what DecodeDecimalBig — the exponent and coefficient of a SIGNED-DECIMAL or UNSIGNED-DECIMAL Value of any coefficient size
package core

import (
	"errors"
	"math/big"
)

// DecodeDecimalBig returns E and C of r (C×10^E), a SIGNED-DECIMAL (two's
// complement coefficient) or UNSIGNED-DECIMAL (magnitude coefficient) Value
// whose coefficient may exceed int64. The coefficient must be minimal.
func DecodeDecimalBig(r Record) (int32, *big.Int, error) {
	if r.Kind != KindValue || len(r.Payload) < 3 {
		return 0, nil, errors.New("geul/core: expected a decimal Value with exponent and coefficient words")
	}
	e := int32(uint32(r.Payload[0])<<16 | uint32(r.Payload[1]))
	switch r.Codec {
	case CodecSignedDecimal:
		c, err := wordsToBigInt(r.Payload[2:])
		return e, c, err
	case CodecUnsignedDecimal:
		c, err := wordsToBigUint(r.Payload[2:])
		return e, c, err
	default:
		return 0, nil, errors.New("geul/core: expected a SIGNED-DECIMAL or UNSIGNED-DECIMAL Value")
	}
}
