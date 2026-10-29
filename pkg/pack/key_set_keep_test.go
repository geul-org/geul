//ff:func feature=pack type=helper control=sequence
//ff:what TestKeySetKeep — the first location kept for a key wins
package pack

import "testing"

func TestKeySetKeep(t *testing.T) {
	s := keySet{}
	if s.keep("a", Loc{Off: 1}) != nil || s.keep("a", Loc{Off: 2}) != nil || s["a"].Off != 1 {
		t.Fatal(s)
	}
}
