//ff:func feature=core type=helper control=selection
//ff:what QuantifierName — the spec name of a 3-bit Reference Quantifier (QQQ)
package core

// QuantifierName returns the name of Quantifier q (GEUL t002 OVERVIEW
// Quantifier): 000 Identified, 001 Generic, 010 Free-choice, 011 Universal,
// 100 Variable, 101-111 Reserved. Only the low 3 bits are read.
func QuantifierName(q uint8) string {
	switch q & 7 {
	case 0:
		return "Identified"
	case 1:
		return "Generic"
	case 2:
		return "Free-choice"
	case 3:
		return "Universal"
	case 4:
		return "Variable"
	default:
		return "Reserved"
	}
}
