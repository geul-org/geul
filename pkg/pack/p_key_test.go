//ff:func feature=pack type=encoder control=sequence
//ff:what TestPKey — big-endian bytes so string order is (Issuer, document, slot) order
package pack

import "testing"

func TestPKey(t *testing.T) {
	k := pKey(1, 0x510000000000002a, 2)
	if len(k) != 16 || k[3] != 1 || k[4] != 0x51 || k[15] != 2 {
		t.Fatalf("%x", k)
	}
	if !(pKey(0, 9, 9) < pKey(1, 0, 0) && pKey(1, 1, 9) < pKey(1, 2, 0) && pKey(1, 2, 0) < pKey(1, 2, 1)) {
		t.Fatal("order")
	}
}
