//ff:func feature=inspect type=helper control=sequence
//ff:what TestNodeRefs — every REF field in layout order; a non-REF claim object and absent fields add nothing
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestNodeRefs(t *testing.T) {
	c := toulmin.TID(3)
	p := toulmin.NodePacket{
		Func:    &toulmin.FuncRef{Alias: 2},
		Check:   &c,
		Claim:   &toulmin.Claim{Subject: 4, Kind: toulmin.ObjTID, ObjTID: 5},
		With:    []toulmin.TID{6},
		Grounds: []toulmin.TID{7},
	}
	if got := nodeRefs(p); !slices.Equal(got, []uint64{2, 3, 4, 5, 6, 7}) {
		t.Fatalf("all: %v", got)
	}
	num := toulmin.NodePacket{Claim: &toulmin.Claim{Subject: 4, Kind: toulmin.ObjNum, ObjNum: 5}}
	if got := nodeRefs(num); !slices.Equal(got, []uint64{4}) {
		t.Fatalf("numeric object: %v", got)
	}
	if got := nodeRefs(toulmin.NodePacket{}); got != nil {
		t.Fatalf("empty: %v", got)
	}
}
