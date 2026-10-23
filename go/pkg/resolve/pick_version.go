//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what pickVersion — the list entry for a version ID, or the latest non-deprecated one for 0
package resolve

// pickVersion finds want in list. For want 0 it returns the highest version that
// is not deprecated. The bool is false when there is no such entry.
func pickVersion(list []VersionEntry, want uint32) (VersionEntry, bool) {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Version == want || (want == 0 && !list[i].Deprecated) {
			return list[i], true
		}
	}
	return VersionEntry{}, false
}
