//ff:func feature=resolve type=helper control=sequence
//ff:what TestHashPath — h/0x + 64 hex digits + .geul
package resolve

import (
	"strings"
	"testing"
)

func TestHashPath(t *testing.T) {
	var h [32]byte
	h[31] = 0xAB
	if p := hashPath(h); p != "h/0x"+strings.Repeat("0", 62)+"ab.geul" {
		t.Fatalf("got %q", p)
	}
}
