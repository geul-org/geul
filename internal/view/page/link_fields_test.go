//ff:func feature=page type=builder control=sequence
//ff:what TestLinkFields — REF and function dst, flags, gate and attack force; a bare link has only src and dst
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestLinkFields(t *testing.T) {
	f := linkFields(toulmin.EdgePacket{Src: 1, DstTID: 2}, nil)
	if len(f) != 2 || f[1].Key != "dst" || f[1].Refs[0].ID != 2 {
		t.Fatalf("bare: %+v", f)
	}
	p := toulmin.EdgePacket{Src: 1, DstFunc: &toulmin.FuncRef{Alias: 3}, Flags: toulmin.EdgeFlagOnce | toulmin.EdgeFlagGate, Gate: &toulmin.Gate{Op: toulmin.GateAtLeast, Percent: 75}}
	f = linkFields(p, nil)
	if len(f) != 4 || f[1].Text != "verb 0x0000" || f[2].Text != "once" || f[3].Text != "≥ 75%" {
		t.Fatalf("full: %+v", f)
	}
	half := toulmin.VFix(5000)
	f = linkFields(toulmin.EdgePacket{Src: 1, DstTID: 2, Flags: toulmin.EdgeFlagForce, Force: &half}, nil)
	if len(f) != 3 || f[2].Key != "force" || f[2].Text != "0.5" {
		t.Fatalf("force: %+v", f)
	}
}
