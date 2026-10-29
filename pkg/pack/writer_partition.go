//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.partition — split a state into sorted P entries, H entries and the NAMED table
package pack

import "sort"

// partition sorts st into the three key spaces.
func (w *writer) partition(st map[string]Loc) ([]kv, []kv, []namedEntry) {
	var ps, hs []kv
	var ns []namedEntry
	for p, l := range st {
		ref := parsePath(p)
		switch ref.space {
		case spaceP:
			ps = append(ps, kv{pKey(w.inum[ref.issuer], ref.doc, ref.slot), l})
		case spaceH:
			hs = append(hs, kv{ref.hash, l})
		default:
			ns = append(ns, namedEntry{p, l})
		}
	}
	sort.Slice(ns, func(i, j int) bool { return ns[i].path < ns[j].path })
	return sortKVs(ps), sortKVs(hs), ns
}
