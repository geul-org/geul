//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadHKeys — 32-byte keys; a short read fails the cursor
package pack

import (
	"strings"
	"testing"
)

func TestReadHKeys(t *testing.T) {
	c := &cursor{b: []byte(strings.Repeat("a", 32) + "b")}
	if k := readHKeys(c, 1); k[0] != strings.Repeat("a", 32) {
		t.Fatal(k)
	}
	if readHKeys(c, 1); !c.bad {
		t.Fatal("short")
	}
}
