//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeGate — round trip plus missing op, reserved bits, bad record, wrong codec and range errors
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeGate(t *testing.T) {
	w, _ := EncodeGate(Gate{Op: GateAtMost, Percent: 100})
	if g, n, err := DecodeGate(w); err != nil || n != len(w) || g.Op != GateAtMost || g.Percent != 100 {
		t.Fatalf("got %+v %d %v", g, n, err)
	}
	if _, _, err := DecodeGate(nil); err == nil {
		t.Fatal("missing operator must fail")
	}
	if _, _, err := DecodeGate([]core.Word{0x0001}); err == nil {
		t.Fatal("reserved bits must fail")
	}
	if _, _, err := DecodeGate([]core.Word{0}); err == nil {
		t.Fatal("missing percent record must fail")
	}
	uv, _ := core.EncodeRecord(core.EncodeUintValue(5))
	if _, _, err := DecodeGate(append([]core.Word{0}, uv...)); err == nil {
		t.Fatal("non-decimal percent must fail")
	}
	short, _ := core.EncodeRecord(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedDecimal, Payload: []core.Word{0, 0}})
	if _, _, err := DecodeGate(append([]core.Word{0}, short...)); err == nil {
		t.Fatal("short decimal must fail")
	}
	big, _ := core.EncodeRecord(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedDecimal, Payload: []core.Word{0, 0, 101}})
	if _, _, err := DecodeGate(append([]core.Word{0}, big...)); err == nil {
		t.Fatal("percent over 100 must fail")
	}
}
