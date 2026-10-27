//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeGate — operator word plus percent Value, and the over-100 error
package toulmin

import "testing"

func TestEncodeGate(t *testing.T) {
	w, err := EncodeGate(Gate{Op: GateAbove, Percent: 75})
	if err != nil || w[0] != 0x4000 || len(w) != 5 || w[4] != 75 {
		t.Fatalf("got %04X %v", w, err)
	}
	if _, err := EncodeGate(Gate{Percent: 101}); err == nil {
		t.Fatal("percent over 100 must fail")
	}
}
