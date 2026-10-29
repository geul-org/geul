//ff:func feature=pack type=writer control=sequence
//ff:what Archive.Pending — whether staged changes wait for Commit
package pack

// Pending reports staged writes or removals.
func (a *Archive) Pending() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return len(a.puts)+len(a.dels) > 0
}
