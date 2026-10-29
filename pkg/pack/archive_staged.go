//ff:func feature=pack type=writer control=sequence
//ff:what Archive.staged — Batch.Get over the staged bytes
package pack

import "io/fs"

// staged returns the staged bytes of name.
func (a *Archive) staged(name string) ([]byte, error) {
	b, ok := a.puts[name]
	if !ok {
		return nil, pathErr("read", name, fs.ErrNotExist)
	}
	return b, nil
}
