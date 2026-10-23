//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestBindDoc — the document of a target-0 EXTERNAL-ISSUER of the list's Issuer; other kinds, codecs, targets and Issuers fail
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestBindDoc(t *testing.T) {
	if d, ok := bindDoc(subjectOf(2, 0), "0x0000"); !ok || d != "0xffff000000000002" {
		t.Fatalf("ok: %q %v", d, ok)
	}
	bad := []core.Record{
		subjectOf(2, 1),
		core.EncodeLocalRef(0),
		{Kind: core.KindValue, Codec: core.RefExternalIssuer1Word << 3},
		{Kind: core.KindReference, Codec: core.RefExternalIssuer1Word << 3, Payload: []core.Word{0, 0xffff, 0, 0, 2}},
	}
	for i, r := range bad {
		if _, ok := bindDoc(r, "0x0000"); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
	if _, ok := bindDoc(subjectOf(2, 0), "0x0001"); ok {
		t.Fatal("other Issuer: want false")
	}
}
