//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what headSources — the Edge Ids of the HEAD's SOURCEs on the whole document
package schema

// headSources replays the HEAD Edges of head (headEdgeRec) and returns the
// Edge Ids of its SOURCEs.
func headSources(head []rec) map[uint64]bool {
	st := &placeState{}
	for _, r := range head {
		headEdgeRec(st, r)
	}
	return st.sources
}
