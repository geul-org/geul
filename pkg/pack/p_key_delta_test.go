//ff:func feature=pack type=encoder control=sequence
//ff:what TestPKeyDelta — raw first key, Issuer step, document step, slot step
package pack

import "testing"

func TestPKeyDelta(t *testing.T) {
	if a, b, c := pKeyDelta(0, 0, 0, 0, 1, 5, 2); a != 1 || b != 5 || c != 2 {
		t.Fatal("first")
	}
	if a, b, c := pKeyDelta(1, 1, 5, 2, 3, 4, 1); a != 2 || b != 4 || c != 1 {
		t.Fatal("issuer")
	}
	if a, b, c := pKeyDelta(1, 1, 5, 2, 1, 6, 0); a != 0 || b != 1 || c != 0 {
		t.Fatal("doc")
	}
	if a, b, c := pKeyDelta(1, 1, 6, 0, 1, 6, 1); a != 0 || b != 0 || c != 1 {
		t.Fatal("slot")
	}
}
