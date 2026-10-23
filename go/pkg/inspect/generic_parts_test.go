//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestGenericParts — an Edge's leading LOCAL id, a Meta's Value, a bare payload and an empty payload
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestGenericParts(t *testing.T) {
	cases := []struct {
		e     Entry
		names []string
	}{
		{Entry{Record: core.Record{Kind: core.KindEdge, Payload: []core.Word{0x00C0, 1, 9}}}, []string{"ID", "payload"}},
		{Entry{Record: core.Record{Kind: core.KindMeta, Payload: []core.Word{0x0882, 1}}, Data: Value{}}, []string{"value"}},
		{Entry{Record: core.Record{Kind: core.KindEdge, Payload: []core.Word{0x0882, 1}}}, []string{"payload"}},
		{Entry{Record: core.Record{Kind: core.KindMeta}}, nil},
	}
	for i, c := range cases {
		var names []string
		for _, p := range genericParts(c.e) {
			names = append(names, p.Name)
		}
		if !slices.Equal(names, c.names) {
			t.Fatalf("case %d: %v", i, names)
		}
	}
}
