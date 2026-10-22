//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderImportDoc — with and without a pinned version; an impossible layout fails
package geulmd

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderImportDoc(t *testing.T) {
	if got, ok := renderImportDoc(core.Record{Payload: []core.Word{3, 1, 0, 0, 0, 2}}, 1); !ok || got != "IMPORT-DOC #3 0x0001 0x0000000000000002" {
		t.Fatalf("plain: %q %v", got, ok)
	}
	pinned := append([]core.Word{3, 1, 0, 0, 0, 2}, make([]core.Word, 16)...)
	if got, ok := renderImportDoc(core.Record{Payload: pinned}, 1); !ok || !strings.Contains(got, " v=0x"+strings.Repeat("0", 64)) {
		t.Fatalf("pinned: %q %v", got, ok)
	}
	if _, ok := renderImportDoc(core.Record{Payload: []core.Word{3, 1, 1, 1, 0, 0, 0, 2}}, 1); ok {
		t.Fatal("3-word issuer: want not ok")
	}
}
