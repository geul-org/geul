//ff:func feature=schema type=rule control=selection
//ff:what checkReference — R1~R4 for one Reference record
package schema

import "fmt"

// checkReference applies the Reference rules: reserved Codec and Quantifier
// (R1), then the layout of the pointer, issuer or declaration forms.
func checkReference(r rec, imports map[uint64]bool) []Fact {
	c, q := r.R.Codec>>3, r.R.Codec&7
	var fs []Fact
	if q >= 5 {
		fs = append(fs, factAt("R1", Warn, r, "Quantifier", "000~100", fmt.Sprintf("%03b (예약)", q)))
	}
	switch {
	case c >= 0xD:
		return append(fs, factAt("R1", Warn, r, "Codec", "0000~1100", fmt.Sprintf("%04b (예약)", c)))
	case c <= 3:
		return append(fs, checkPointer(r, c, imports)...)
	case c <= 7:
		return append(fs, checkIssuerRef(r, c)...)
	default:
		return append(fs, checkTarget(r, r.R.Payload)...)
	}
}
