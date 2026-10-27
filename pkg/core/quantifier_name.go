//ff:func feature=core type=helper control=selection
//ff:what QuantifierName — the spec name of a Reference's Quantifier (QQQ); "" for Codec 0000, whose low 3 bits are a slot marker number
package core

// QuantifierName returns the name of the Quantifier in the 7-bit Reference
// codec c (CCCC << 3 | QQQ; GEUL reference Quantifier): 000 Identified, 001
// Generic, 010 Free-choice, 011 Universal, 100 Variable, 101-111 Reserved.
// Exception: under Codec 0000 (slot markers) the low 3 bits are a marker
// number, not a Quantifier, so it returns "" (see SlotMarkerName).
func QuantifierName(c uint8) string {
	switch {
	case c>>3&0xF == 0:
		return ""
	case c&7 == 0:
		return "Identified"
	case c&7 == 1:
		return "Generic"
	case c&7 == 2:
		return "Free-choice"
	case c&7 == 3:
		return "Universal"
	case c&7 == 4:
		return "Variable"
	default:
		return "Reserved"
	}
}
