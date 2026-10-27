//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what DecodeFraction — the numerator and denominator records of a FRACTION Value, checking their codecs and that nothing is left over
package core

import (
	"errors"
	"fmt"
)

// DecodeFraction reads the two records of a FRACTION Value r
// (value/numeric/FRACTION): a SIGNED-INTEGER or SIGNED-DECIMAL numerator and
// an UNSIGNED-INTEGER or UNSIGNED-DECIMAL denominator, the payload ending
// right after the second. Minimal words, the exponent of a DECIMAL component
// and a zero denominator are schema V10 (and FractionRat).
func DecodeFraction(r Record) (Fraction, error) {
	if r.Kind != KindValue || r.Codec != CodecFraction {
		return Fraction{}, fmt.Errorf("geul/core: expected a FRACTION Value, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	var rs [2]Record
	ws := r.Payload
	for i := range rs {
		rec, n, err := DecodeRecord(ws)
		if err != nil {
			return Fraction{}, fmt.Errorf("geul/core: FRACTION record %d: %w", i+1, err)
		}
		rs[i], ws = rec, ws[n:]
	}
	if len(ws) != 0 {
		return Fraction{}, errors.New("geul/core: FRACTION has words after the denominator")
	}
	if err := fractionForm(rs[0], rs[1]); err != nil {
		return Fraction{}, err
	}
	return Fraction{Numerator: rs[0], Denominator: rs[1]}, nil
}
