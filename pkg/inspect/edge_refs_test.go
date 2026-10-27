//ff:func feature=inspect type=helper control=sequence
//ff:what TestEdgeRefs — src then dst; a function dst gives its ALIAS
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestEdgeRefs(t *testing.T) {
	if got := edgeRefs(toulmin.EdgePacket{Src: 1, DstTID: 2}); !slices.Equal(got, []uint64{1, 2}) {
		t.Fatalf("ref: %v", got)
	}
	if got := edgeRefs(toulmin.EdgePacket{Src: 1, DstTID: 9, DstFunc: &toulmin.FuncRef{Alias: 3}}); !slices.Equal(got, []uint64{1, 3}) {
		t.Fatalf("func: %v", got)
	}
}
