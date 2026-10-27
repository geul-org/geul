//ff:func feature=schema type=rule control=selection
//ff:what checkReference — R1~R4 for one Reference record
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkReference applies the Reference rules: reserved Codec and Quantifier
// (R1), then the layout of the pointer, issuer, issuer/document pointer or
// declaration forms.
func checkReference(r rec, imports map[uint64]bool) []Fact {
	c, q := r.R.Codec>>3, r.R.Codec&7
	var fs []Fact
	if q >= 5 {
		fs = append(fs, factAt("R1", Warn, r, "Quantifier", "000~100", fmt.Sprintf("%03b (예약)", q)))
	}
	switch {
	case c == 0:
		return append(fs, factAt("R1", Warn, r, "Codec", "0001~1111", "0000 (예약)"))
	case c == core.RefIssuerPointer || c == core.RefDocPointer:
		return append(fs, checkIssuerDocPointer(r, c)...)
	case c >= core.RefExternalIssuer1Word:
		return append(fs, checkIssuerRef(r, c)...)
	case c >= core.RefExternalPointer1Word:
		return append(fs, checkPointer(r, c, imports)...)
	default:
		return append(fs, checkTarget(r, r.R.Payload)...)
	}
}
