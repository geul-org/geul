//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what deletedIn — whether a version list has the version and marks it deleted (TOMBSTONE)
package resolve

// deletedIn reports whether list has version v and it is deleted.
func deletedIn(list []VersionEntry, v uint32) bool {
	for _, e := range list {
		if e.Version == v && e.Deleted {
			return true
		}
	}
	return false
}
