//ff:func feature=schema type=rule control=sequence
//ff:what TestIsVersionList — document ID 0, another ID, no DOC-NAME
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestIsVersionList(t *testing.T) {
	name := func(doc core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocName, Payload: []core.Word{0, 0, 0, 0, doc}})
	}
	if !isVersionList([]rec{name(0)}) || isVersionList([]rec{name(5)}) || isVersionList(nil) {
		t.Fatal("wrong")
	}
}
