//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestControlLine — PROSE, SECTION and case controls
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestControlLine(t *testing.T) {
	s := toulmin.Span{Off: 1, Len: 2}
	if controlLine("PROSE", toulmin.ControlPacket{ID: 1, Code: toulmin.ControlProse, Span: &s}) != "PROSE #1 @1+2" ||
		controlLine("SECTION", toulmin.ControlPacket{ID: 2, Code: toulmin.ControlSection, Arg: 5}) != "SECTION #2 5" ||
		controlLine("CASE-END", toulmin.ControlPacket{ID: 3, Code: toulmin.ControlCaseEnd, Arg: 4}) != "CASE-END #3 #4" {
		t.Fatal("control line mismatch")
	}
}
