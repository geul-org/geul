//ff:func feature=notation type=encoder control=sequence
//ff:what parseFraction — FRACTION from two component tokens: numerator, denominator, each kept as written
package notation

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseFraction reads FRACTION(n, d): the numerator signed and the
// denominator unsigned (parseFractionPart), encodes them and checks the
// result as a FRACTION (core.DecodeFraction). It does not reduce.
func parseFraction(args []string) (core.Record, error) {
	if len(args) != 2 {
		return core.Record{}, errors.New("geul/notation: FRACTION takes numerator, denominator")
	}
	n, err1 := parseFractionPart(args[0], true)
	d, err2 := parseFractionPart(args[1], false)
	if err := errors.Join(err1, err2); err != nil {
		return core.Record{}, err
	}
	r, _ := core.EncodeFraction(core.Fraction{Numerator: n, Denominator: d})
	_, err := core.DecodeFraction(r)
	return r, err
}
