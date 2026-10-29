//ff:func feature=pack type=writer control=selection
//ff:what writer.entryOf — the inventory entry of a written record (kind 1, 2 or 3; flags bit0 when it serves h/)
package pack

// entryOf describes p for the current pack's inventory.
func (w *writer) entryOf(p *put) invEntry {
	e := invEntry{enc: p.loc.Enc, n: uint32(p.loc.Len), off: p.loc.Off}
	if p.share != "" {
		e.flags = flagHash
	}
	switch p.ref.space {
	case spaceP:
		e.kind, e.ref, e.doc, e.slot = invP, w.inum[p.ref.issuer], p.ref.doc, p.ref.slot
	case spaceH:
		e.kind = invH
	default:
		e.kind, e.ref = invNamed, uint32(len(w.inv.paths))
		w.inv.paths = append(w.inv.paths, p.path)
	}
	return e
}
