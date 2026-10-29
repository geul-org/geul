//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.Paths — every live path of the archive (all sets, staged changes applied), sorted
package pack

// Paths lists the archive's paths.
func (a *Archive) Paths() ([]string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	idx, err := a.index()
	return append([]string(nil), idx...), err
}
