//ff:func feature=schema type=rule control=selection
//ff:what checkReference — R1~R5 for one Reference record (Codec 0000 by checkCodecZero)
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkReference applies the Reference rules: the reserved Codec 0000, whose
// low bits are a slot marker number and no Quantifier (checkCodecZero), the
// reserved Quantifier (R1), then the layout of the pointer, issuer,
// issuer/document pointer or declaration forms.
func checkReference(r rec, imports map[uint64]bool) []Fact {
	c, q := r.R.Codec>>3, r.R.Codec&7
	if c == 0 {
		return checkCodecZero(r)
	}
	var fs []Fact
	if q >= 5 {
		fs = append(fs, factAt("R1", Warn, r, "Quantifier", "000~100", fmt.Sprintf("%03b (예약)", q)))
	}
	switch {
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
