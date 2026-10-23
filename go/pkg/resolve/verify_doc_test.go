//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what TestVerifyDoc — matching hash and DOC-ID pass; odd bytes, malformed, wrong hash and wrong name fail with ErrVerify
package resolve

import (
	"errors"
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestVerifyDoc(t *testing.T) {
	b := sampleDoc(1, "x")
	ws, _ := core.ReadWords(b)
	h, _ := core.ContentHash(ws)
	v := VersionEntry{Version: 3, Hash: h}
	if n, err := verifyDoc(b, sampleKey(1, 0), v); err != nil || n.Version != 3 || n.Hash != h || len(n.Doc.Body) != 1 {
		t.Fatalf("ok: %+v %v", n, err)
	}
	bad := map[string]func() error{
		"odd":       func() error { _, err := verifyDoc([]byte{1}, sampleKey(1, 0), v); return err },
		"malformed": func() error { _, err := verifyDoc([]byte{0, 0}, sampleKey(1, 0), v); return err },
		"hash":      func() error { _, err := verifyDoc(b, sampleKey(1, 0), VersionEntry{Version: 3}); return err },
		"name":      func() error { _, err := verifyDoc(b, sampleKey(2, 0), v); return err },
	}
	for name, f := range bad {
		if err := f(); !errors.Is(err, ErrVerify) {
			t.Fatalf("%s: %v", name, err)
		}
	}
}
