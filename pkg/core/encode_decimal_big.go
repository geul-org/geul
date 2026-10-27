//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDecimalBig — a SIGNED-DECIMAL or UNSIGNED-DECIMAL Value C×10^E of any coefficient size
package core

import (
	"errors"
	"math/big"
)

// EncodeDecimalBig builds the decimal Value c×10^e for a coefficient of any
// size: the 2-word exponent, then c in minimal words — two's complement for
// SIGNED-DECIMAL (signed), magnitude for UNSIGNED-DECIMAL. C and E are kept as
// given (quantum). A negative c under UNSIGNED-DECIMAL is an error.
func EncodeDecimalBig(signed bool, e int32, c *big.Int) (Record, error) {
	if signed {
		return EncodeSignedDecimalBig(e, c), nil
	}
	if c.Sign() < 0 {
		return Record{}, errors.New("geul/core: UNSIGNED-DECIMAL coefficient is negative")
	}
	payload := append([]Word{Word(uint32(e) >> 16), Word(uint32(e))}, bigUintWords(c)...)
	return Record{Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: payload}, nil
}
