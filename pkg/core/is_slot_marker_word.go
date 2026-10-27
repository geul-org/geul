//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarkerWord — whether a WORD is a slot marker TAG WORD or exact CODEC target (0x0000~0x0007)
package core

// IsSlotMarkerWord reports whether w is 0x0000~0x0007: as a TAG WORD, a
// one-word Reference of Codec 0000 (a slot marker); as a CODEC target, the
// exact target naming that marker. One bit test, no codebook.
func IsSlotMarkerWord(w Word) bool {
	return w&0xFFF8 == 0
}
