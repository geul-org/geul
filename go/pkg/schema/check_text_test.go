//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckText — UTF-8 and BOM for text, ASCII for URI, language tags, EDTF and GEDCOM-DATE (spaces inside), plain-day EDTF fails V6
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
		{0x40, "\xff", "V5"},
		{0x40, "\xef\xbb\xbfx", "V5"},
		{0x44, "http://x", ""},
		{0x44, "", "V5"},
		{0x46, "ko KR", "V5"},
		{0x24, "1974", ""},
		{0x24, "197X", ""},
		{0x24, "1952-03-11?", ""},
		{0x24, "", "V5"},
		{0x24, "1952-03-11", "V6"},
		{0x24, "-0043-03-15", "V6"},
		{0x26, "JULIAN 14 OCT 1066", ""},
		{0x26, "JULIAN 1300 ", "V5"},
		{0x26, "", "V5"},
	}
	for i, c := range cases {
		if got := rulesOf(checkText(rec{}, c.even, []byte(c.b))); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
