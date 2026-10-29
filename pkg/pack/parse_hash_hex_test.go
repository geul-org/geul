//ff:func feature=pack type=decoder control=sequence
//ff:what TestParseHashHex — 64 lowercase hex digits only
package pack

import (
	"strings"
	"testing"
)

func TestParseHashHex(t *testing.T) {
	if h, ok := parseHashHex(strings.Repeat("0f", 32)); !ok || h[31] != 0x0f {
		t.Fatal(h)
	}
	if _, ok := parseHashHex("0f"); ok {
		t.Fatal("short")
	}
}
