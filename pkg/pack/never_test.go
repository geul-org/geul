//ff:func feature=pack type=helper control=sequence
//ff:what never — an fsFuncs that never fails
package pack

// never returns a pass-through fileSys.
func never() *fsFuncs {
	return &fsFuncs{mkdir: -1, create: -1, rename: -1}
}
