//ff:func feature=pack type=helper control=sequence
//ff:what fsFuncs.Remove — the real remove
package pack

// Remove deletes name.
func (f *fsFuncs) Remove(name string) error {
	return osFS{}.Remove(name)
}
