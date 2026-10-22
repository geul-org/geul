//ff:func feature=inspect type=decoder control=sequence
//ff:what TestParseImportDoc — with and without a pinned version; an impossible layout fails
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseImportDoc(t *testing.T) {
	d, ok := parseImportDoc(core.Record{Payload: []core.Word{3, 1, 0, 0, 0, 2}}, 1)
	if !ok || d.ID != 3 || len(d.Issuer) != 1 || d.Doc[3] != 2 || d.Version != nil {
		t.Fatalf("plain: %+v %v", d, ok)
	}
	pinned := append([]core.Word{0, 3, 1, 0, 0, 0, 2}, make([]core.Word, 16)...)
	if d, ok := parseImportDoc(core.Record{Payload: pinned}, 2); !ok || d.ID != 3 || len(d.Version) != 16 {
		t.Fatalf("pinned: %+v %v", d, ok)
	}
	if _, ok := parseImportDoc(core.Record{Payload: []core.Word{3, 1, 1, 1, 0, 0, 0, 2}}, 1); ok {
		t.Fatal("3-word issuer: want not ok")
	}
}
