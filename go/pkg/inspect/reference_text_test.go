//ff:func feature=inspect type=decoder control=sequence
//ff:what TestReferenceText — a Wikidata item as Qn, another Reference as codec name and hex
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestReferenceText(t *testing.T) {
	if got := referenceText(core.Record{Kind: core.KindReference, Codec: core.RefWikidataQID << 3, Payload: []core.Word{0x000C, 0xA250}}); got != "Q827984" {
		t.Fatalf("qid: %q", got)
	}
	if got := referenceText(core.EncodeLocalRef(3)); got != "LOCAL 0x0003" {
		t.Fatalf("local: %q", got)
	}
}
