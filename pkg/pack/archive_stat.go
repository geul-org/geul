//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.Stat — fs.StatFS: the info of a file or directory
package pack

import "io/fs"

// Stat describes name.
func (a *Archive) Stat(name string) (fs.FileInfo, error) {
	f, err := a.Open(name)
	if err != nil {
		return nil, err
	}
	return f.Stat()
}
