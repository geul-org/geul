//ff:type feature=resolve type=resolver
//ff:what Store — the bytes under an archive's `.well-known/geul/` as an fs.FS that can also be written (directory tree or pack sets)
package resolve

import "io/fs"

// Store holds an archive's paths (relative to `.well-known/geul/`, fs.FS
// names: "." is the top). Reads follow io/fs: a missing path wraps
// fs.ErrNotExist. WriteFile creates what it needs; Remove of a missing path
// wraps fs.ErrNotExist. A directory tree (the default of Dir) and pack sets
// (geul pkg/pack Archive) implement it.
type Store interface {
	fs.ReadDirFS
	fs.ReadFileFS
	WriteFile(name string, b []byte) error
	Remove(name string) error
}
