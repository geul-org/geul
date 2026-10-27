//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeVFixValue — read a SIGNED-DECIMAL Value record back into a VFix in [-1.0, 1.0]
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// DecodeVFixValue reads a SIGNED-DECIMAL Value record as a VFix. Any exponent is
// accepted as long as the value is exact at four decimal places and lies in the
// verdict range [-1.0, +1.0].
func DecodeVFixValue(r core.Record) (VFix, error) {
	if r.Codec != core.CodecSignedDecimal {
		return 0, errors.New("geul/toulmin: verdict/qualifier must be a SIGNED-DECIMAL Value")
	}
	e, c, err := core.DecodeDecimalValue(r)
	if err != nil {
		return 0, err
	}
	v, err := core.ScaleDecimal(e, c, -4, vfixScale)
	if err != nil {
		return 0, err
	}
	return VFix(v), nil
}
