//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.recordBytes — the original bytes of every record this generation writes
package pack

// recordBytes sums the sizes of the records (shared h/ puts excluded).
func (w *writer) recordBytes() int64 {
	var n int64
	for _, p := range w.records() {
		n += int64(p.n)
	}
	return n
}
