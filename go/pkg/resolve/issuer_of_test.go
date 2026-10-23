//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestIssuerOf — (Issuer, 0, 0) of each width; a document, a target, other codecs, kinds and short payloads are not an Issuer
package resolve

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestIssuerOf(t *testing.T) {
	ref := func(c uint8, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c << 3, Payload: p}
	}
	if s, ok := issuerOf(ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 0, 0)); !ok || s != "0x0000" {
		t.Fatalf("1 word: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.RefExternalIssuer2Word, 0, 0x1234, 0, 0, 0, 0, 0)); !ok || s != "0x00001234" {
		t.Fatalf("2 words: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.RefExternalIssuer16Word, append(make([]core.Word, 16), 0, 0, 0, 0, 0)...)); !ok || s != "0x"+strings.Repeat("0", 64) {
		t.Fatalf("16 words: %q %v", s, ok)
	}
	bad := []core.Record{
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 5, 0),
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 0, 7),
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 0),
		ref(core.RefLocal, 0),
		{Kind: core.KindValue, Codec: core.RefExternalIssuer1Word << 3, Payload: []core.Word{0, 0, 0, 0, 0, 0}},
	}
	for i, r := range bad {
		if _, ok := issuerOf(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
