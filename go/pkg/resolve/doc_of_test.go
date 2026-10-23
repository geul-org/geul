//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestDocOf — (Issuer, document, 0) is a document; the Issuer itself, a target, other codecs, kinds and short payloads are not
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDocOf(t *testing.T) {
	ref := func(c uint8, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c << 3, Payload: p}
	}
	if d, ok := docOf(ref(core.RefExternalIssuer2Word, 0, 0x20, 0xffff, 0, 0, 3, 0)); !ok || d != (DocKey{Issuer: "0x00000020", Doc: "0xffff000000000003"}) {
		t.Fatalf("got %+v %v", d, ok)
	}
	for i, r := range []core.Record{
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 0, 0),
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 1, 7),
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 1),
		ref(core.RefLocal, 0),
		{Kind: core.KindValue, Codec: core.RefExternalIssuer1Word << 3, Payload: []core.Word{0, 0, 0, 0, 1, 0}},
	} {
		if _, ok := docOf(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
