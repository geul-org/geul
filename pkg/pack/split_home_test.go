//ff:func feature=pack type=writer control=sequence
//ff:what TestSplitHome — 0x0000 first, else the first Issuer in Dir order, else the top set
package pack

import "testing"

func TestSplitHome(t *testing.T) {
	if splitHome(map[string]bool{"0x00010000": true, "0x0000": true}) != "0x0000" || splitHome(map[string]bool{"0x00010000": true, "0x0001": true}) != "0x0001" || splitHome(nil) != "" {
		t.Fatal("splitHome")
	}
}
