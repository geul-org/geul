//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckImportDoc — too short and a non-minimal IMPORT ID fail M3
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckImportDoc(t *testing.T) {
	imp := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc1Word, Payload: p})
	}
	if checkImportDoc(imp(1, 1, 0, 0, 0, 2, 0, 0)) != nil || rulesOf(checkImportDoc(imp(1, 0, 0, 0, 2, 0, 0))) != "M3" || rulesOf(checkImportDoc(imp(0, 1, 1, 0, 0, 0, 2, 0, 0))) != "M3" {
		t.Fatal("wrong")
	}
}
