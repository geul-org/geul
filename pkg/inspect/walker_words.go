//ff:func feature=inspect type=decoder control=sequence
//ff:what walker.words — take n plain words as a Part
package inspect

// words takes n plain words at the cursor as one part. Fewer than n remaining
// words marks the walk bad.
func (w *walker) words(name string, n int) {
	if w.bad || len(w.ws)-w.pos < n {
		w.bad = true
		return
	}
	w.out = append(w.out, Part{Name: name, Words: w.ws[w.pos : w.pos+n]})
	w.pos += n
}
