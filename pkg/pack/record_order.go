//ff:func feature=pack type=writer control=selection
//ff:what recordOrder — a put's sort key for data packs: P records by key, then h/-only records by hash, then named by path
package pack

// recordOrder sorts records as §6 step 2 asks (versions.geul, slot 0, before
// the versions of its document).
func recordOrder(p *put, inum map[string]uint32) string {
	switch p.ref.space {
	case spaceP:
		return "\x01" + pKey(inum[p.ref.issuer], p.ref.doc, p.ref.slot)
	case spaceH:
		return "\x02" + p.ref.hash
	default:
		return "\x03" + p.path
	}
}
