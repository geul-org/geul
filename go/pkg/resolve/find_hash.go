//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what findHash — the list entry already holding a content hash
package resolve

// findHash returns the entry of list whose content hash is h.
func findHash(list []VersionEntry, h [32]byte) (VersionEntry, bool) {
	for _, e := range list {
		if e.Hash == h {
			return e, true
		}
	}
	return VersionEntry{}, false
}
