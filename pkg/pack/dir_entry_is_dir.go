//ff:func feature=pack type=helper control=sequence
//ff:what dirEntry.IsDir — whether the entry is a directory
package pack

// IsDir reports a directory.
func (e dirEntry) IsDir() bool { return e.dir }
