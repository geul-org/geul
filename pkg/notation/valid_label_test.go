//ff:func feature=notation type=rule control=iteration dimension=1
//ff:what TestValidLabel — digits, letters, _ and - are labels; empty and other characters are not
package notation

import "testing"

func TestValidLabel(t *testing.T) {
	for s, want := range map[string]bool{"7": true, "edu": true, "a_b-2": true, "": false, "a.b": false, "가": false, "a b": false} {
		if validLabel(s) != want {
			t.Errorf("%q: want %v", s, want)
		}
	}
}
