//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseImportDoc — latest and pinned versions, a 2-word IMPORT ID; a non-IMPORT codec, a short payload and a non-minimal ID fail
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseImportDoc(t *testing.T) {
	latest := []core.Word{3, 1, 0, 0, 0, 2, 0, 0}
	d, ok := parseImportDoc(core.Record{Codec: core.MetaImportDoc1Word, Payload: latest})
	if !ok || d.ID != 3 || len(d.Issuer) != 1 || d.Doc[3] != 2 || d.Version != 0 {
		t.Fatalf("latest: %+v %v", d, ok)
	}
	pinned := []core.Word{1, 3, 1, 0, 0, 0, 2, 1, 9}
	if d, ok := parseImportDoc(core.Record{Codec: core.MetaImportDoc1Word, Payload: pinned}); !ok || d.ID != 0x10003 || d.Version != 0x10009 {
		t.Fatalf("pinned: %+v %v", d, ok)
	}
	bad := []core.Record{
		{Codec: core.MetaDocName, Payload: latest},
		{Codec: core.MetaImportDoc4Word, Payload: latest},
		{Codec: core.MetaImportDoc1Word, Payload: append([]core.Word{0}, latest...)},
	}
	for i, r := range bad {
		if _, ok := parseImportDoc(r); ok {
			t.Fatalf("bad %d: want not ok", i)
		}
	}
}
