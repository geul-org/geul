//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestIssuerOf — ISSUER-POINTER of each width; the old EXTERNAL-ISSUER form, a document, other kinds and bad widths are not an Issuer
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
	if s, ok := issuerOf(ref(core.RefIssuerPointer, 0)); !ok || s != "0x0000" {
		t.Fatalf("1 word: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.RefIssuerPointer, 0, 0x1234)); !ok || s != "0x00001234" {
		t.Fatalf("2 words: %q %v", s, ok)
	}
	if s, ok := issuerOf(ref(core.RefIssuerPointer, make([]core.Word, 16)...)); !ok || s != "0x"+strings.Repeat("0", 64) {
		t.Fatalf("16 words: %q %v", s, ok)
	}
	bad := []core.Record{
		ref(core.RefExternalIssuer1Word, 0, 0, 0, 0, 0, 0),
		ref(core.RefDocPointer, 0, 0xffff, 0, 0, 1, 0, 0),
		ref(core.RefIssuerPointer, 0, 0, 0),
		{Kind: core.KindValue, Codec: core.RefIssuerPointer << 3, Payload: []core.Word{0}},
	}
	for i, r := range bad {
		if _, ok := issuerOf(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
