//ff:func feature=pack type=helper control=sequence
//ff:what TestIssuerLess — width first, then hex
package pack

import "testing"

func TestIssuerLess(t *testing.T) {
	if !issuerLess("0xffff", "0x00000000") || !issuerLess("0x0000", "0x0001") || issuerLess("0x0001", "0x0000") {
		t.Fatal("issuerLess")
	}
}
