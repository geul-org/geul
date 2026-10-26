//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what newestLive — the highest version of a list that is neither deprecated nor deleted
package resolve

// newestLive returns the highest version in es that is not deleted or
// deprecated; false when there is none.
func newestLive(es []VersionEntry) (VersionEntry, bool) {
	var best VersionEntry
	found := false
	for _, e := range es {
		if !e.Deleted && !e.Deprecated && (!found || e.Version > best.Version) {
			best, found = e, true
		}
	}
	return best, found
}
