//ff:func feature=inspect type=decoder control=sequence
//ff:what walker.span — take a SPAN (offset and length at the tried SPAN width) as two Parts
package inspect

// span takes a SPAN at the tried width: name+".off" then name+".len".
func (w *walker) span(name string) {
	w.words(name+".off", w.sw)
	w.words(name+".len", w.sw)
}
