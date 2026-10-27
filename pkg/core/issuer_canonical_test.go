//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerCanonical — the shortest width holding the value passes; wider forms of a smaller value and other widths fail
package core

import "testing"

func TestIssuerCanonical(t *testing.T) {
	high16 := make([]Word, 16)
	high16[11] = 1
	for i, c := range []struct {
		issuer []Word
		want   bool
	}{
		{[]Word{0}, true},
		{[]Word{0x1234}, true},
		{[]Word{1, 0}, true},
		{[]Word{0, 0x1234}, false},
		{[]Word{0, 1, 0, 0}, true},
		{[]Word{0, 0, 0xFFFF, 0xFFFF}, false},
		{high16, true},
		{make([]Word, 16), false},
		{[]Word{1, 2, 3}, false},
		{nil, false},
	} {
		if IssuerCanonical(c.issuer) != c.want {
			t.Fatalf("%d: %04X want %v", i, c.issuer, c.want)
		}
	}
}
