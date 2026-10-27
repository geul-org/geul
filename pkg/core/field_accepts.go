//ff:func feature=core type=helper control=selection
//ff:what FieldAccepts — whether a record fits a record field's CODEC list: slot markers only when named exactly, nothing else of Codec 0000, any other record when the list is empty
package core

import "slices"

// FieldAccepts reports whether r may fill a record field whose CODEC targets
// are ts (GEUL edge/CODEC). A Reference of the reserved Codec 0000 is accepted
// only as a slot marker (one word, IsSlotMarker) whose exact WORD the list
// names: a field without CODECs ("any record") and wildcards never take one.
// Any other record fits an empty list, or a list with a target it matches
// (CodecMatches).
func FieldAccepts(ts []Word, r Record) bool {
	marker := r.Kind == KindReference && r.Codec>>3 == 0
	switch {
	case marker && !IsSlotMarker(r):
		return false
	case len(ts) == 0:
		return !marker
	default:
		return slices.ContainsFunc(ts, func(t Word) bool { return CodecMatches(t, r.Kind, r.Codec) })
	}
}
