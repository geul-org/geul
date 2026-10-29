//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.store — the Dir's Store: the one set, else the directory tree under Root
package resolve

// store returns d.Store or a dirStore over `<Root>/.well-known/geul/`.
func (d Dir) store() Store {
	if d.Store != nil {
		return d.Store
	}
	return dirStore{root: d.file("")}
}
