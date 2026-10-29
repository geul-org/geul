//ff:func feature=pack type=writer control=sequence
//ff:what writer.closeObject — append the tail, patch the header, hash, rename to o/<sha256>, add to OBJECTS
package pack

import "path/filepath"

// closeObject finishes w.cur; tail is a data pack's inventory (nil for an index object).
func (w *writer) closeObject(tail []byte) error {
	ow := w.cur
	w.cur = nil
	h := objectHeader{kind: ow.kind, count: ow.count}
	if ow.kind == kindData {
		h.invOff, h.invLen = ow.off, uint64(len(tail))
	}
	ow.write(tail)
	h.size = ow.off
	if ow.err == nil {
		_, ow.err = ow.f.WriteAt(encodeObjectHeader(h), 0)
	}
	var sum [32]byte
	if ow.err == nil {
		sum, ow.err = sumFile(ow.f, h.size)
	}
	if err := ow.f.Close(); ow.err == nil {
		ow.err = err
	}
	o := object{kind: ow.kind, size: h.size, hash: sum}
	if ow.err == nil {
		ow.err = w.fs.Rename(ow.f.Name(), filepath.Join(w.out, filepath.FromSlash(objectName(o))))
	}
	if ow.err != nil {
		w.fs.Remove(ow.f.Name())
		return ow.err
	}
	w.objs = append(w.objs, o)
	w.res.Objects = append(w.res.Objects, objectName(o))
	w.res.Bytes += int64(h.size)
	return nil
}
