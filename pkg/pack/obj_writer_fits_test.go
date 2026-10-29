//ff:func feature=pack type=helper control=sequence
//ff:what TestObjWriterFits — within the limit, or anything into an empty object
package pack

import "testing"

func TestObjWriterFits(t *testing.T) {
	w := &objWriter{off: 64}
	if !w.fits(1000, 100) {
		t.Fatal("empty object takes an oversized item")
	}
	w.count = 1
	if !w.fits(36, 100) || w.fits(37, 100) {
		t.Fatal("limit")
	}
}
