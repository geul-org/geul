//ff:func feature=pack type=writer control=sequence
//ff:what writer.startObject — open a temp file for the next object, numbered by its coming place in OBJECTS
package pack

import "path/filepath"

// startObject opens a new object of kind after a 64-byte header placeholder.
func (w *writer) startObject(kind byte) error {
	f, err := w.fs.CreateTemp(filepath.Join(w.out, "o"))
	if err != nil {
		return err
	}
	w.cur, w.curNum = &objWriter{f: f, kind: kind}, uint32(len(w.objs))
	w.cur.write(make([]byte, headerSize))
	return nil
}
