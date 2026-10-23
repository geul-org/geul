//ff:func feature=resolve type=rule control=sequence
//ff:what TestEncodeVersionList — HEAD SPEC-VERSION, optional PARENT and RECORDED-AT, then the records
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
	"testing"
)

func TestEncodeVersionList(t *testing.T) {
	rec := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Doc: []core.Word{0, 0, 0, 1}, Version: 1})
	first, _ := inspect.Decode(encodeVersionList([]core.Record{rec}, nil, 5))
	if len(first.Head) != 2 || first.Head[0].Name != "SPEC-VERSION" || first.Head[1].Name != "RECORDED-AT" || len(first.Body) != 1 {
		t.Fatalf("first: %+v", first.Head)
	}
	parent := [32]byte{1}
	next, _ := inspect.Decode(encodeVersionList([]core.Record{rec, rec}, &parent, 6))
	if len(next.Head) != 3 || next.Head[1].Name != "PARENT" || len(next.Body) != 2 {
		t.Fatalf("next: %+v", next.Head)
	}
}
