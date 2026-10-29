//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.carryDicts — the previous DICTS renumbered; none on Fresh (records are re-encoded by this codec)
package pack

// carryDicts keeps dictionaries that carried records may still need.
func (w *writer) carryDicts() []dictEntry {
	if w.o.Fresh {
		return nil
	}
	out := make([]dictEntry, 0, len(w.prev.m.dicts))
	for _, d := range w.prev.m.dicts {
		d.loc = w.mapLoc(d.loc)
		out = append(out, d)
	}
	return out
}
