//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.Name — the base name
package pack

// Name is the last path element.
func (f fileInfo) Name() string { return f.name }
