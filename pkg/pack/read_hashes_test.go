//ff:func feature=pack type=decoder control=sequence
//ff:what TestReadHashes — 32-byte hashes; an impossible count fails the cursor
package pack

import "testing"

func TestReadHashes(t *testing.T) {
	c := &cursor{b: append([]byte{1, 0, 0, 0, 0, 0, 0, 0}, make([]byte, 32)...)}
	if h := readHashes(c); len(h) != 1 || len(h[0]) != 32 || c.bad {
		t.Fatal(h)
	}
	c = &cursor{b: []byte{2, 0, 0, 0, 0, 0, 0, 0}}
	if readHashes(c); !c.bad {
		t.Fatal("count")
	}
}
