//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestDocOf — DOC-POINTER(Issuer, document, 0) is a document; a version, the old EXTERNAL-ISSUER form, other codecs and kinds are not
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDocOf(t *testing.T) {
	ref := func(c uint8, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c << 3, Payload: p}
	}
	if d, ok := docOf(ref(core.RefDocPointer, 1, 0x20, 0xffff, 0, 0, 3, 0, 0)); !ok || d != (DocKey{Issuer: "0x00010020", Doc: "0xffff000000000003"}) {
		t.Fatalf("got %+v %v", d, ok)
	}
	for i, r := range []core.Record{
		ref(core.RefDocPointer, 0, 0xffff, 0, 0, 3, 0, 1),
		ref(core.RefExternalIssuer1Word, 0, 0xffff, 0, 0, 3, 0),
		ref(core.RefIssuerPointer, 0),
		{Kind: core.KindValue, Codec: core.RefDocPointer << 3, Payload: []core.Word{0, 0xffff, 0, 0, 3, 0, 0}},
	} {
		if _, ok := docOf(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
