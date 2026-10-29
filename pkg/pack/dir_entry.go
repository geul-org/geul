//ff:type feature=pack type=model
//ff:what dirEntry — an fs.DirEntry of an Archive listing: a name and whether it is a directory
package pack

// dirEntry is one ReadDir entry.
type dirEntry struct {
	name string
	dir  bool
	size int64
	a    *Archive // set on listed files: Info reads the size through it
	path string
}
