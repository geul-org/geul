//ff:func feature=pack type=resolver control=selection
//ff:what invPath — the path a P or named inventory entry (record or deletion) names
package pack

// invPath resolves an entry's path through the pack's Issuer or path table.
func invPath(inv inventory, e invEntry) (string, error) {
	switch e.kind {
	case invP, invPDel:
		if int(e.ref) >= len(inv.issuers) {
			return "", corrupt("inventory Issuer number")
		}
		return pPath(inv.issuers[e.ref], e.doc, e.slot), nil
	default:
		if int(e.ref) >= len(inv.paths) {
			return "", corrupt("inventory path number")
		}
		return inv.paths[e.ref], nil
	}
}
