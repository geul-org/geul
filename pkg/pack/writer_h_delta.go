//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.hDelta — the new H layer's entries: changed h/ paths and deletion markers for removed ones
package pack

// hDelta lists this generation's H changes, sorted.
func (w *writer) hDelta() []kv {
	var out []kv
	for _, p := range w.puts {
		if p.ref.space == spaceH {
			out = append(out, kv{p.ref.hash, p.finalLoc()})
		}
	}
	for _, d := range w.dels {
		if ref := parsePath(d); ref.space == spaceH {
			out = append(out, kv{ref.hash, Loc{Enc: encGone}})
		}
	}
	return sortKVs(out)
}
