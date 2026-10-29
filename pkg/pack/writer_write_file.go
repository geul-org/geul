//ff:func feature=pack type=writer control=sequence
//ff:what writer.writeFile — write a small file atomically: temp file in the same directory, then rename
package pack

import (
	"path/filepath"
)

// writeFile publishes b at name (slash path under out).
func (w *writer) writeFile(name string, b []byte) error {
	dst := filepath.Join(w.out, filepath.FromSlash(name))
	f, err := w.fs.CreateTemp(filepath.Dir(dst))
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = w.fs.Rename(f.Name(), dst)
	}
	if err != nil {
		w.fs.Remove(f.Name())
	}
	return err
}
