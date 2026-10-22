//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestLinkLine — REF and FuncRef dsts, once only where not implied, strike, ordered and gate
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestLinkLine(t *testing.T) {
	g := toulmin.Gate{Op: toulmin.GateAbove, Percent: 75}
	all := toulmin.EdgeFlagOnce | toulmin.EdgeFlagStrike | toulmin.EdgeFlagOrdered | toulmin.EdgeFlagGate
	if got := linkLine("UNDO", toulmin.EdgePacket{ID: 5, Src: 1, DstFunc: &toulmin.FuncRef{Alias: 2}, Flags: all, Gate: &g}); got != "UNDO #5 #1 -> #2/0x0000 once strike ordered gate=above/75 @0+0" {
		t.Fatalf("undo: %q", got)
	}
	if got := linkLine("DO-ONCE", toulmin.EdgePacket{ID: 6, Src: 1, DstFunc: &toulmin.FuncRef{}, Flags: toulmin.EdgeFlagOnce}); got != "DO-ONCE #6 #1 -> #0/0x0000 @0+0" {
		t.Fatalf("do-once: %q", got)
	}
	if got := linkLine("ATTACKS", toulmin.EdgePacket{ID: 7, Src: 1, DstTID: 2}); got != "ATTACKS #7 #1 -> #2 @0+0" {
		t.Fatalf("attacks: %q", got)
	}
}
