//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what entryOf — the list entry of a version number
package resolve

// entryOf returns the entry of version v in es; false when the list has none.
func entryOf(es []VersionEntry, v uint32) (VersionEntry, bool) {
	for _, e := range es {
		if e.Version == v {
			return e, true
		}
	}
	return VersionEntry{}, false
}
