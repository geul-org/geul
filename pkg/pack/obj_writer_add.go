//ff:func feature=pack type=helper control=sequence
//ff:what objWriter.add — append one record or block and return its offset
package pack

// add appends b as one counted item.
func (w *objWriter) add(b []byte) uint64 {
	off := w.off
	w.write(b)
	w.count++
	return off
}
