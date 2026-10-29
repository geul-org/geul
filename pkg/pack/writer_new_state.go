//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.newState — the new generation's live paths: previous state renumbered, minus deletions, plus puts
package pack

// newState is S(M) of the generation being written (§7.1).
func (w *writer) newState() map[string]Loc {
	st := make(map[string]Loc, len(w.base)+len(w.puts))
	for p, l := range w.base {
		st[p] = w.mapLoc(l)
	}
	for _, d := range w.dels {
		delete(st, d)
	}
	for _, p := range w.puts {
		st[p.path] = p.finalLoc()
	}
	return st
}
