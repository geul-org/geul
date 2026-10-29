//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Tree — the archive's Store: the given one, or the directory tree under Root
package resolve

// Tree is the Store a Dir reads and writes; viewers list and read it.
func (d Dir) Tree() Store {
	return d.store()
}
