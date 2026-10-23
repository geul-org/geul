//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestPlace — every placement case of D2~D4
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestPlace(t *testing.T) {
	meta := func(c uint8) rec { return rec{Total: 1, R: core.Record{Kind: core.KindMeta, Codec: c}} }
	val := recOf(core.EncodeUintValue(1))
	st, parts := &placeState{}, &docParts{}
	if fs := place(st, parts, val); fs != nil || !st.inBody || len(parts.Body) != 1 {
		t.Fatalf("implicit body: %+v", fs)
	}
	st, parts = &placeState{}, &docParts{}
	steps := []struct {
		r    rec
		want string
	}{
		{meta(core.MetaSpecVersion), ""},
		{val, "D3"},
		{meta(core.MetaBody), ""},
		{meta(core.MetaBody), "D3"},
		{meta(core.MetaSpecVersion), "D3"},
		{meta(core.MetaDoc1Word), "D4"},
		{val, ""},
		{meta(core.MetaEnd), ""},
		{val, "D2"},
	}
	for i, s := range steps {
		if got := rulesOf(place(st, parts, s.r)); got != s.want {
			t.Fatalf("step %d: %q", i, got)
		}
	}
	if len(parts.Head) != 1 || len(parts.Body) != 3 {
		t.Fatalf("parts: %d %d", len(parts.Head), len(parts.Body))
	}
}
