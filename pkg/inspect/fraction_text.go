//ff:func feature=inspect type=decoder control=sequence
//ff:what fractionText — a FRACTION Value as numerator/denominator, each component exactly as written (2/3, 23.45/1256.1)
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// fractionText decodes a FRACTION (core.DecodeFraction) and writes it as
// n/d with each component by numberText: not reduced, digits kept.
func fractionText(r core.Record) (string, error) {
	f, err := core.DecodeFraction(r)
	if err != nil {
		return "", err
	}
	n, err := numberText(f.Numerator)
	if err != nil {
		return "", err
	}
	d, err := numberText(f.Denominator)
	return n + "/" + d, err
}
