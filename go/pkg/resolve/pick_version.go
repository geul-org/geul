//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what pickVersion — the list entry for a version ID, or the latest live one for 0; deleted versions never
package resolve

// pickVersion finds want in list. For want 0 it returns the highest version that
// is neither deprecated nor deleted. A deleted version is never returned. The
// bool is false when there is no such entry.
func pickVersion(list []VersionEntry, want uint32) (VersionEntry, bool) {
	for i := len(list) - 1; i >= 0; i-- {
		e := list[i]
		if !e.Deleted && (e.Version == want || (want == 0 && !e.Deprecated)) {
			return e, true
		}
	}
	return VersionEntry{}, false
}
