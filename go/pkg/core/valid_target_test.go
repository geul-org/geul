//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestValidTarget — exact and wildcard targets pass; prefix bits, other LLLL and wildcards with nothing to match fail
package core

import "testing"

func TestValidTarget(t *testing.T) {
	for w, want := range map[Word]bool{
		0x084A: true, 0x0008: true, 0x0780: true, 0x0F80: true, 0x07A0: true, 0x0781: true,
		0x2008: false, 0x0088: false, 0x0F84: false, 0x07A1: false,
	} {
		if got := validTarget(w); got != want {
			t.Fatalf("%04X: %v", w, got)
		}
	}
}
