//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what findDict — the DICTS entry of a dictionary ID
package pack

// findDict looks id up in ds.
func findDict(ds []dictEntry, id uint32) (dictEntry, bool) {
	for _, d := range ds {
		if d.id == id {
			return d, true
		}
	}
	return dictEntry{}, false
}
