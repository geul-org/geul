//ff:func feature=page type=helper control=sequence
//ff:what quantityText — a QUANTITY as amount, unit and [lower, upper]: "652 Q828224 [651, 653]", "1.20"
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// quantityText shows a QUANTITY: the exact amount, then its unit when it has
// one, then the bounds as [lower, upper] when either is given (a missing one
// as ?).
func quantityText(v inspect.Value) string {
	s := v.Amount
	if v.Unit != "" {
		s += " " + v.Unit
	}
	if v.Lower == "" && v.Upper == "" {
		return s
	}
	or := func(x string) string {
		if x == "" {
			return "?"
		}
		return x
	}
	return s + " [" + or(v.Lower) + ", " + or(v.Upper) + "]"
}
