//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestValidName — slugs pass; empty, dashed ends, an ID prefix and other characters fail
package resolve

import "testing"

func TestValidName(t *testing.T) {
	for s, want := range map[string]bool{"q42": true, "delta-base": true, "": false, "-a": false, "a-": false, "0x12": false, "A": false, "a.b": false} {
		if validName(s) != want {
			t.Errorf("%q: want %v", s, want)
		}
	}
}
