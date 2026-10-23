//ff:func feature=inspect type=decoder control=sequence
//ff:what TestApplySpec — declared entries get their fields; undeclared ones none
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
	entries := []Entry{{Record: r}, {Record: core.EncodeUintValue(1)}, {Record: core.Record{Kind: core.KindEdge, Codec: 0x05}}}
	applySpec(entries, s)
	if entries[0].Spec == nil || entries[1].Spec == nil || entries[2].Spec != nil {
		t.Fatalf("got %+v", entries)
	}
}
