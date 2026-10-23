//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocument — document-wide rules, version list rules only for document ID 0
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckDocument(t *testing.T) {
	name := func(doc core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocName, Payload: []core.Word{0, 0, 0, 0, doc}})
	}
	body := []rec{recOf(edgeOf(0, 1)), recOf(edgeOf(0, 1))}
	if fs := checkDocument(docParts{Head: []rec{name(5)}, Body: body}); rulesOf(fs) != "E2" {
		t.Fatalf("plain: %+v", fs)
	}
	if fs := checkDocument(docParts{Head: []rec{name(0)}, Body: body}); rulesOf(fs) != "E2,P1,P1" {
		t.Fatalf("list: %+v", fs)
	}
}
