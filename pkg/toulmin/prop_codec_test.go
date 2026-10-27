//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestPropCodec — every relation round-trips; do+once is DO-ONCE; provenance and unknown fail
package toulmin

import "testing"

func TestPropCodec(t *testing.T) {
	for _, p := range props {
		ref, err := PropCodec(p, false)
		if err != nil {
			t.Fatalf("0x%04X: %v", uint16(p), err)
		}
		got, once, ok := PropFromCodec(ref)
		if !ok || got != p || once {
			t.Fatalf("0x%04X: round trip got 0x%04X once=%v ok=%v", uint16(p), uint16(got), once, ok)
		}
	}
	if ref, _ := PropCodec(PropDo, true); ref.Codec != CodecDoOnce {
		t.Fatalf("do once: %+v", ref)
	}
	if _, err := PropCodec(PropProvenance, false); err == nil {
		t.Fatal("provenance must fail")
	}
	if _, err := PropCodec(0x0FFF, false); err == nil {
		t.Fatal("unknown PropCode must fail")
	}
}
