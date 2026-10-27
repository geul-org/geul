//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestFmtRef — a format number lives under ISSUER-1WORD
package toulmin

import "testing"

func TestFmtRef(t *testing.T) {
	if r := fmtRef(FormatText); r.Codec != 0x7C || r.Format != FormatText {
		t.Fatalf("got %+v", r)
	}
}
