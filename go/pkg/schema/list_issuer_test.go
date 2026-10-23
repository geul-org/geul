//ff:func feature=schema type=rule control=sequence
//ff:what TestListIssuer — the DOC-ID's Issuer; empty without DOC-ID
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestListIssuer(t *testing.T) {
	id := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{0, 0x1234, 0, 0, 0, 0}})
	if got := listIssuer([]rec{recOf(core.Record{Kind: core.KindMeta, Codec: 0x10, Payload: []core.Word{2}}), id}); got != "0x00001234" {
		t.Fatalf("got %q", got)
	}
	if listIssuer(nil) != "" {
		t.Fatal("none")
	}
}
