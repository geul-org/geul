//ff:func feature=schema type=rule control=sequence
//ff:what TestMetaInstant — the first matching Meta's instant; absent and malformed
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestMetaInstant(t *testing.T) {
	from := recOf(metaOf(core.MetaEffectiveFrom, core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{0, 1}}))
	if tm, _, ok := metaInstant([]rec{from}, core.MetaEffectiveFrom); !ok || tm.Unix() != 86400 {
		t.Fatalf("ok: %v %v", tm, ok)
	}
	if _, _, ok := metaInstant([]rec{from}, core.MetaEffectiveUntil); ok {
		t.Fatal("absent: want false")
	}
	bad := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaEffectiveFrom})
	if _, _, ok := metaInstant([]rec{bad}, core.MetaEffectiveFrom); ok {
		t.Fatal("malformed: want false")
	}
}
