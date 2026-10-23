//ff:func feature=core type=helper control=sequence
//ff:what SplitTarget — the Kind and Codec of a DECLARE target word
package core

// SplitTarget reads the Kind and Codec bits of a target word (TargetWord).
func SplitTarget(w Word) (Kind, uint8) {
	return Kind(w >> 11 & 3), uint8(w & 0x7F)
}
