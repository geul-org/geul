//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestGedcomASCII — inner spaces pass; leading or trailing spaces, control bytes and non-ASCII fail
package schema

import "testing"

func TestGedcomASCII(t *testing.T) {
	for s, want := range map[string]bool{
		"JULIAN 14 OCT 1066": true, "1300": true,
		" JULIAN 1300": false, "JULIAN 1300 ": false, "JULIAN\t1300": false, "JULIAN 1300\xc3": false,
	} {
		if gedcomASCII([]byte(s)) != want {
			t.Fatalf("%q: want %v", s, want)
		}
	}
}
