//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what blameStep — fold one version into the blame table: new Edge Ids are added, different content is changed
package delta

// blameStep records version v of snapshot s in seen: an Edge Id not seen is
// added (and changed) at v, one whose content differs from the last seen is
// changed at v.
func blameStep(seen map[uint64]*Line, s snapshot, v uint32) {
	for i, id := range s.ids {
		l, ok := seen[id]
		switch {
		case !ok:
			seen[id] = &Line{ID: id, Record: s.body[i], Added: v, Changed: v}
		case !sameRecord(l.Record, s.body[i]):
			l.Record, l.Changed = s.body[i], v
		}
	}
}
