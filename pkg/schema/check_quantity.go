//ff:func feature=schema type=rule control=sequence
//ff:what checkQuantity — V9: a QUANTITY holds a SIGNED-DECIMAL amount, a Reference or Absent unit, SIGNED-DECIMAL or Absent bounds, all minimal, and nothing more
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkQuantity applies V9 (value/numeric/QUANTITY): the four records decode
// in their forms (core.DecodeQuantity) and the amount and each present bound
// have a minimal coefficient (core.DecodeSignedDecimalBig).
func checkQuantity(r rec) []Fact {
	q, err := core.DecodeQuantity(r.R)
	if err != nil {
		return []Fact{factAt("V9", Error, r, "Payload", "amount·unit·lower·upper 레코드", err.Error())}
	}
	if err := quantityDecimals(q); err != nil {
		return []Fact{factAt("V9", Error, r, "Payload", "최소 WORD coefficient", err.Error())}
	}
	return nil
}
