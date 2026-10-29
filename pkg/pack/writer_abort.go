//ff:func feature=pack type=writer control=sequence
//ff:what writer.abort — drop the temp file of an object left open by a failed write
package pack

// abort closes and removes w.cur, if any.
func (w *writer) abort() {
	if w.cur == nil {
		return
	}
	w.cur.f.Close()
	w.fs.Remove(w.cur.f.Name())
	w.cur = nil
}
