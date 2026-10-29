//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.isDir — whether any live path lies under name (the top always is a directory)
package pack

// isDir reports a directory; listing errors count as no.
func (a *Archive) isDir(name string) bool {
	a.mu.Lock()
	idx, err := a.index()
	a.mu.Unlock()
	return name == "." || err == nil && hasPrefix(idx, dirPrefix(name))
}
