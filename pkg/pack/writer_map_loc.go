//ff:func feature=pack type=writer control=sequence
//ff:what writer.mapLoc — a carried location renumbered for the rebuilt OBJECTS (unchanged when OBJECTS only grows)
package pack

// mapLoc applies w.remap to l.
func (w *writer) mapLoc(l Loc) Loc {
	if w.remap != nil {
		l.Obj = w.remap[l.Obj]
	}
	return l
}
