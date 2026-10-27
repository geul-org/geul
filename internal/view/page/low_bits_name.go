//ff:func feature=page type=helper control=selection
//ff:what lowBitsName — what a Reference TAG's low 3 bits say: the Quantifier, or under Codec 0000 the slot marker number
package page

import "github.com/geul-org/gogeul/pkg/core"

// lowBitsName names the low 3 bits of the 7-bit Reference codec c: the
// Quantifier (core.QuantifierName), except under the reserved Codec 0000,
// where they are a slot marker number, not a Quantifier ("표지 번호", with
// "예약" for 3~7).
func lowBitsName(c uint8) string {
	switch {
	case c>>3 != 0:
		return core.QuantifierName(c)
	case core.SlotMarkerName(c) == "":
		return "표지 번호 (예약)"
	default:
		return "표지 번호"
	}
}
