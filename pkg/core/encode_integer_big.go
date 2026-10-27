//ff:func feature=core type=encoder control=sequence
//ff:what EncodeIntegerBig — a SIGNED-INTEGER or UNSIGNED-INTEGER Value of any size in minimal words
package core

import (
	"errors"
	"math/big"
)

// EncodeIntegerBig builds the integer Value of c: SIGNED-INTEGER (signed, two's
// complement) or UNSIGNED-INTEGER (magnitude), in the fewest words. A negative
// c under UNSIGNED-INTEGER is an error.
func EncodeIntegerBig(signed bool, c *big.Int) (Record, error) {
	if signed {
		return Record{Kind: KindValue, Codec: CodecSignedInteger, Payload: bigIntWords(c)}, nil
	}
	if c.Sign() < 0 {
		return Record{}, errors.New("geul/core: UNSIGNED-INTEGER is negative")
	}
	return Record{Kind: KindValue, Codec: CodecUnsignedInteger, Payload: bigUintWords(c)}, nil
}
