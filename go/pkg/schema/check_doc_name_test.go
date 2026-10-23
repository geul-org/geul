//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocName — Issuer widths 1/2/4/16 pass, 3 fails M2
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckDocName(t *testing.T) {
	dn := func(n int) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocName, Payload: make([]core.Word, n)})
	}
	if checkDocName(dn(5)) != nil || checkDocName(dn(20)) != nil || rulesOf(checkDocName(dn(7))) != "M2" {
		t.Fatal("wrong")
	}
}
