//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderGeneration — the manifest's generation
package pack

import "testing"

func TestReaderGeneration(t *testing.T) {
	if g := newReader(memStore{}, Options{}, &manifest{gen: 4}, "").Generation(); g != 4 {
		t.Fatal(g)
	}
}
