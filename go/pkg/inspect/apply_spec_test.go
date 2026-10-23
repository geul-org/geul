//ff:func feature=inspect type=decoder control=sequence
//ff:what TestApplySpec — declared entries get fields and an Edge Id; undeclared, non-Edge and already-identified entries keep theirs
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestApplySpec(t *testing.T) {
	r := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 3, Doc: []core.Word{0, 0, 0, 1}, Version: 1})
	s := tombstoneSpec()
	s[core.TargetWord(core.KindValue, core.CodecUnsignedInteger)] = []codecspec.Field{{Name: "n", Width: 1}}
	entries := []Entry{{Record: r}, {Record: r, ID: 9}, {Record: core.EncodeUintValue(1)}, {Record: core.Record{Kind: core.KindEdge, Codec: 0x05}}}
	applySpec(entries, s)
	if entries[0].ID != 3 || entries[0].Spec == nil || entries[1].ID != 9 || entries[2].Spec == nil || entries[2].ID != 0 || entries[3].Spec != nil {
		t.Fatalf("got %+v", entries)
	}
}
