//ff:func feature=pack type=helper control=sequence
//ff:what fileInfo.Sys — nothing underlying
package pack

// Sys is nil.
func (f fileInfo) Sys() any { return nil }
