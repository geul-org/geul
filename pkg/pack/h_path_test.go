//ff:func feature=pack type=encoder control=sequence
//ff:what TestHPath — h/0x + 64 hex + .geul, the inverse of parsePath for H
package pack

import (
	"strings"
	"testing"
)

func TestHPath(t *testing.T) {
	k := strings.Repeat("\x01", 32)
	p := hPath(k)
	if p != "h/0x"+strings.Repeat("01", 32)+".geul" || parsePath(p).hash != k {
		t.Fatal(p)
	}
}
