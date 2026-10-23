//ff:func feature=codecspec type=helper control=sequence
//ff:what TestKeyOf — Kind·Codec, a Reference's Quantifier cleared
package codecspec

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestKeyOf(t *testing.T) {
	if k := keyOf(core.Record{Kind: core.KindEdge, Codec: core.EdgePublication}); k != 0x1018 {
		t.Fatalf("edge: %04X", k)
	}
	if k := keyOf(core.EncodeLocalRef(1)); k != core.TargetWord(core.KindReference, core.CodecLocalIdentified&^7) {
		t.Fatalf("reference: %04X", k)
	}
}
