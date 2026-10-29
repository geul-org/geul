//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseZstd — no options, a level, nodict, both; out-of-range, repeated and unknown options fail
package pack

import "testing"

func TestParseZstd(t *testing.T) {
	cases := map[string][2]int{"": {0, 0}, "19": {19, 0}, "nodict": {0, -1}, "nodict 3": {3, -1}}
	for in, want := range cases {
		c, err := parseZstd("s", splitFields(in))
		if z := c.(*Zstd); err != nil || z.Level != want[0] || z.DictSize != want[1] {
			t.Fatal(in, err)
		}
	}
	for _, in := range []string{"23", "0", "3 4", "nodict nodict", "fast"} {
		if _, err := parseZstd("s", splitFields(in)); err == nil {
			t.Fatal(in)
		}
	}
}
