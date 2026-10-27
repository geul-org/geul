//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocID — an EXTERNAL-ISSUER-*(Issuer, document, 0) passes; a raw payload fails M2
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckDocID(t *testing.T) {
	ok, _ := core.EncodeDocID([]core.Word{1, 1}, []core.Word{0, 0, 0, 5})
	raw := core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: make([]core.Word, 5)}
	if checkDocID(recOf(ok)) != nil || rulesOf(checkDocID(recOf(raw))) != "M2" {
		t.Fatal("wrong")
	}
}
