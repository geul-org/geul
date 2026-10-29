//ff:func feature=pack type=helper control=sequence
//ff:what TestSampleOrder — by SHA-256, on a copy
package pack

import "testing"

func TestSampleOrder(t *testing.T) {
	a, b := &put{path: "a", sum: [32]byte{9}}, &put{path: "b", sum: [32]byte{1}}
	in := []*put{a, b}
	out := sampleOrder(in)
	if out[0] != b || out[1] != a || in[0] != a {
		t.Fatal(out)
	}
}
