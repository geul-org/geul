//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what latestStored — the highest version of a list that is not deleted, deprecated or not
package resolve

// latestStored returns the highest version in es that is not deleted (its
// file is still kept); false when there is none.
func latestStored(es []VersionEntry) (VersionEntry, bool) {
	for i := len(es) - 1; i >= 0; i-- {
		if !es[i].Deleted {
			return es[i], true
		}
	}
	return VersionEntry{}, false
}
