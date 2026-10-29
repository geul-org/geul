//ff:func feature=pack type=helper control=sequence
//ff:what objWriter.fits — whether n more bytes stay within the size limit (an empty object takes anything)
package pack

// fits reports whether adding n bytes keeps the object within limit; an
// object with no items accepts one oversized item (§6 step 4).
func (w *objWriter) fits(n int, limit int64) bool {
	return w.count == 0 || w.off+uint64(n) <= uint64(limit)
}
