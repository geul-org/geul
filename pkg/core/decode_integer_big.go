//ff:func feature=core type=decoder control=selection
//ff:what DecodeIntegerBig — the value of a SIGNED-INTEGER or UNSIGNED-INTEGER Value of any size, in minimal words
package core

import (
	"fmt"
	"math/big"
)

// DecodeIntegerBig returns the value of r, a SIGNED-INTEGER (two's
// complement) or UNSIGNED-INTEGER (magnitude) Value record of any size. The
// payload must be minimal.
func DecodeIntegerBig(r Record) (*big.Int, error) {
	switch {
	case r.Kind == KindValue && r.Codec == CodecSignedInteger:
		return wordsToBigInt(r.Payload)
	case r.Kind == KindValue && r.Codec == CodecUnsignedInteger:
		return wordsToBigUint(r.Payload)
	default:
		return nil, fmt.Errorf("geul/core: expected an INTEGER Value, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
}
