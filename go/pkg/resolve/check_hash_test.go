//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestCheckHash — the right hash returns the document; a wrong one, odd bytes and a non-document fail
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCheckHash(t *testing.T) {
	b := edgeDoc(1, "a")
	ws, _ := core.ReadWords(b)
	h, _ := core.ContentHash(ws)
	if _, err := checkHash(b, h); err != nil {
		t.Fatal(err)
	}
	for _, x := range [][]byte{{1}, {0, 5}} {
		if _, err := checkHash(x, h); err == nil {
			t.Fatal("bad bytes: want error")
		}
	}
	if _, err := checkHash(b, [32]byte{1}); err == nil {
		t.Fatal("mismatch: want error")
	}
}
