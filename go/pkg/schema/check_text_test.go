//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckText — UTF-8 validity and BOM, empty and non-ASCII URI or tag
package schema

import (
	"testing"
)

func TestCheckText(t *testing.T) {
	cases := []struct {
		even uint8
		b    string
		want string
	}{
		{0x40, "가", ""},
		{0x42, "\xff", "V5"},
		{0x40, "\xef\xbb\xbfx", "V5"},
		{0x44, "http://x", ""},
		{0x44, "", "V5"},
		{0x46, "ko KR", "V5"},
	}
	for i, c := range cases {
		if got := rulesOf(checkText(rec{}, c.even, []byte(c.b))); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
