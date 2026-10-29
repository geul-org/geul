//ff:func feature=pack type=helper control=sequence
//ff:what dirEntry.Name — the entry name
package pack

// Name is the last path element.
func (e dirEntry) Name() string { return e.name }
