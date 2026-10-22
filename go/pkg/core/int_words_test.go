//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestIntWords — minimal two's-complement forms around the sign-word boundaries
package core

import "testing"

func TestIntWords(t *testing.T) {
	cases := []struct {
		c    int64
		want []Word
	}{
		{0, []Word{0}}, {-1, []Word{0xFFFF}}, {0x7FFF, []Word{0x7FFF}},
		{0x8000, []Word{0, 0x8000}}, {-0x8000, []Word{0x8000}}, {-0x8001, []Word{0xFFFF, 0x7FFF}},
	}
	for _, c := range cases {
		got := intWords(c.c)
		if len(got) != len(c.want) || got[0] != c.want[0] || got[len(got)-1] != c.want[len(c.want)-1] {
			t.Fatalf("%d: got %04X want %04X", c.c, got, c.want)
		}
	}
}
