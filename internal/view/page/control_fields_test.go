//ff:func feature=page type=builder control=sequence
//ff:what TestControlFields — named and unnamed sections, case links, PROSE without fields
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestControlFields(t *testing.T) {
	if f := controlFields(toulmin.ControlPacket{Code: toulmin.ControlSection, Arg: 4}, nil); f[0].Text != "Rules" {
		t.Fatalf("named: %+v", f)
	}
	if f := controlFields(toulmin.ControlPacket{Code: toulmin.ControlSection, Arg: 9}, nil); f[0].Text != "9" {
		t.Fatalf("unnamed: %+v", f)
	}
	if f := controlFields(toulmin.ControlPacket{Code: toulmin.ControlCaseBegin, Arg: 3}, map[uint64]string{3: "c"}); f[0].Refs[0] != (Ref{ID: 3, Label: "c"}) {
		t.Fatalf("case: %+v", f)
	}
	if f := controlFields(toulmin.ControlPacket{Code: toulmin.ControlProse}, nil); f != nil {
		t.Fatalf("prose: %+v", f)
	}
}
