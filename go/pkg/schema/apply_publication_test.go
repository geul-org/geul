//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestApplyPublication — sequence, gap, repeated hash, deprecation of published and unpublished versions
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestApplyPublication(t *testing.T) {
	st := map[string]*docVersions{}
	doc := []core.Word{0, 0, 0, 5}
	pub := func(v uint32, h byte) core.Publication {
		return core.Publication{Codec: core.EdgePublication, Issuer: []core.Word{0}, Doc: doc, Version: v, Hash: [32]byte{h}}
	}
	steps := []struct {
		p    core.Publication
		want string
	}{
		{pub(1, 1), ""},
		{core.Publication{Codec: core.EdgeDeprecation, Issuer: []core.Word{0}, Doc: doc, Version: 1}, ""},
		{core.Publication{Codec: core.EdgeTombstone, Issuer: []core.Word{0}, Doc: doc, Version: 4}, "P4"},
		{pub(3, 2), "P3"},
		{pub(4, 1), "P3"},
		{pub(5, 3), ""},
	}
	for i, s := range steps {
		if got := rulesOf(applyPublication(st, rec{}, s.p)); got != s.want {
			t.Fatalf("step %d: %q", i, got)
		}
	}
}
