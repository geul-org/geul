//ff:func feature=schema type=rule control=sequence
//ff:what checkQuantity — V9: a QUANTITY holds a numeric amount, a Reference or Absent unit, numeric or Absent bounds, each canonical (finite BINARY-FLOAT, valid FRACTION), and nothing more
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkQuantity applies V9 (value/numeric/QUANTITY): the four records decode
// in their forms (core.DecodeQuantity) and the amount and each present bound
// are canonical (quantityNumbers: minimal words, a finite BINARY-FLOAT, a
// valid FRACTION).
func checkQuantity(r rec) []Fact {
	q, err := core.DecodeQuantity(r.R)
	if err != nil {
		return []Fact{factAt("V9", Error, r, "Payload", "amount·unit·lower·upper 레코드", err.Error())}
	}
	if p := quantityNumbers(q); p != "" {
		return []Fact{factAt("V9", Error, r, "Payload", "정규 수치 레코드(최소 WORD, 유한 BINARY-FLOAT, 유효 FRACTION)", p)}
	}
	return nil
}
