//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocID — Issuer widths 1/2/4/16 pass, 3 fails M2
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckDocID(t *testing.T) {
	dn := func(n int) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: make([]core.Word, n)})
	}
	if checkDocID(dn(5)) != nil || checkDocID(dn(20)) != nil || rulesOf(checkDocID(dn(7))) != "M2" {
		t.Fatal("wrong")
	}
}
