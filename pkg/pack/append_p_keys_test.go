//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendPKeys — a document's list and version code as (0, 0, 1) after a document step (§4.7)
package pack

import (
	"bytes"
	"testing"
)

func TestAppendPKeys(t *testing.T) {
	keys := []string{pKey(1, 8, 0), pKey(1, 8, 1), pKey(1, 9, 0)}
	want := []byte{1, 0, 0, 8, 0, 1, 0, 1, 0}
	if b := appendPKeys(nil, keys); !bytes.Equal(b, want) {
		t.Fatalf("%x", b)
	}
}
