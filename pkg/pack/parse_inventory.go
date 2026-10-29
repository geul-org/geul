//ff:func feature=pack type=decoder control=sequence
//ff:what parseInventory — a decoded inventory; trailing bytes and bad tables are ErrCorrupt
package pack

// parseInventory decodes §4.5 after zstd unframing.
func parseInventory(b []byte) (inventory, error) {
	c := &cursor{b: b}
	iss, err := parseIssuers(c)
	if err != nil {
		return inventory{}, err
	}
	inv := inventory{issuers: iss, paths: readPaths(c)}
	inv.entries = readInvEntries(c)
	inv.deleted = readHashes(c)
	if err := c.end("inventory"); err != nil {
		return inventory{}, err
	}
	return inv, nil
}
