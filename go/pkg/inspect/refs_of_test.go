//ff:func feature=inspect type=helper control=sequence
//ff:what TestRefsOf — case controls reference their case; NODE and LINK their REFs; others nothing
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestRefsOf(t *testing.T) {
	if got := refsOf(toulmin.ControlPacket{Code: toulmin.ControlCaseEnd, Arg: 4}); !slices.Equal(got, []uint64{4}) {
		t.Fatalf("case end: %v", got)
	}
	if got := refsOf(toulmin.ControlPacket{Code: toulmin.ControlSection, Arg: 4}); got != nil {
		t.Fatalf("section: %v", got)
	}
	c := toulmin.TID(9)
	if got := refsOf(toulmin.NodePacket{Check: &c}); !slices.Equal(got, []uint64{9}) {
		t.Fatalf("node: %v", got)
	}
	if got := refsOf(toulmin.EdgePacket{Src: 1, DstTID: 2}); !slices.Equal(got, []uint64{1, 2}) {
		t.Fatalf("link: %v", got)
	}
	if got := refsOf(Value{}); got != nil {
		t.Fatalf("value: %v", got)
	}
}
