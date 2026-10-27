//ff:func feature=inspect type=decoder control=sequence
//ff:what quantity — the amount, unit and bounds of a QUANTITY Value as text: exact numbers, Qn unit, empty when left out
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// quantity reads a QUANTITY Value (core.DecodeQuantity) as text: the amount
// and each present bound as exact numbers (numberText: integer, decimal with
// its quantum, BINARY-FLOAT, FRACTION n/d), the unit by referenceText; a
// left-out unit or bound is empty.
func quantity(r core.Record) (amount, unit, lower, upper string, err error) {
	q, err := core.DecodeQuantity(r)
	if err != nil {
		return "", "", "", "", err
	}
	text := func(d core.Record) (string, error) {
		if core.IsAbsent(d) {
			return "", nil
		}
		return numberText(d)
	}
	if !core.IsAbsent(q.Unit) {
		unit = referenceText(q.Unit)
	}
	amount, err = text(q.Amount)
	if err == nil {
		lower, err = text(q.Lower)
	}
	if err == nil {
		upper, err = text(q.Upper)
	}
	return amount, unit, lower, upper, err
}
