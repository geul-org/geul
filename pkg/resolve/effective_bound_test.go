//ff:func feature=resolve type=decoder control=sequence
//ff:what TestEffectiveBound — present and absent bounds; a malformed bound fails
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestEffectiveBound(t *testing.T) {
	doc, _ := inspect.Decode(sampleEffectiveDoc(5, "x", [2]uint32{uint32(core.MetaEffectiveFrom), 200}))
	if tm, ok, err := effectiveBound(doc, core.MetaEffectiveFrom); err != nil || !ok || !tm.Equal(day(200)) {
		t.Fatalf("from: %v %v %v", tm, ok, err)
	}
	if _, ok, err := effectiveBound(doc, core.MetaEffectiveUntil); err != nil || ok {
		t.Fatalf("until: %v %v", ok, err)
	}
	bad := inspect.Document{Head: []inspect.Entry{{Record: core.Record{Kind: core.KindMeta, Codec: core.MetaEffectiveFrom}}}}
	if _, ok, err := effectiveBound(bad, core.MetaEffectiveFrom); err == nil || ok {
		t.Fatalf("bad: %v %v", ok, err)
	}
}
