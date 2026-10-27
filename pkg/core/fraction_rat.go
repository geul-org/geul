//ff:func feature=core type=decoder control=sequence
//ff:what FractionRat — the exact value of a FRACTION Value, (C1×10^E1)/(C2×10^E2); a zero denominator fails
package core

import (
	"errors"
	"math/big"
)

// FractionRat decodes a FRACTION Value (DecodeFraction) and returns
// numerator / denominator exactly (value/numeric/FRACTION 비교): never through
// floating point. A denominator whose value is 0 is an error.
func FractionRat(r Record) (*big.Rat, error) {
	f, err := DecodeFraction(r)
	if err != nil {
		return nil, err
	}
	n, err := NumberRat(f.Numerator)
	if err != nil {
		return nil, err
	}
	d, err := NumberRat(f.Denominator)
	if err != nil {
		return nil, err
	}
	if d.Sign() == 0 {
		return nil, errors.New("geul/core: FRACTION denominator is 0")
	}
	return n.Quo(n, d), nil
}
