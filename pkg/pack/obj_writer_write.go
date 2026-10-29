//ff:func feature=pack type=helper control=sequence
//ff:what objWriter.write — append bytes, keeping the first error
package pack

// write appends b and advances off.
func (w *objWriter) write(b []byte) {
	if w.err == nil {
		_, w.err = w.f.Write(b)
	}
	w.off += uint64(len(b))
}
