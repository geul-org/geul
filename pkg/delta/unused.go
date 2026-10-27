//ff:func feature=delta type=helper control=iteration dimension=1
//ff:what LocalIDs.Unused — the previous Edge Ids no statement of the new version took, in previous order (the diff's DELTA-REMOVE)
package delta

// Unused returns the Edge Ids of the previous version that no statement of
// the new version took, in their previous order: the statements removed
// (DELTA-REMOVE of a diff).
func (l *LocalIDs) Unused() []uint64 {
	var out []uint64
	for _, id := range l.prev {
		if !l.used[id] {
			out = append(out, id)
		}
	}
	return out
}
