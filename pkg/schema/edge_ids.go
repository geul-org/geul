//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what edgeIDs — the set of Edge Ids the BODY declares
package schema

// edgeIDs collects the Edge Id of every BODY record that has one (edgeID).
func edgeIDs(body []rec) map[uint64]bool {
	ids := map[uint64]bool{}
	for _, r := range body {
		if id, ok := edgeID(r); ok {
			ids[id] = true
		}
	}
	return ids
}
