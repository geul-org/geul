//ff:func feature=pack type=helper control=sequence
//ff:what TestPutFinalLoc — its own location, or the sharing record's
package pack

import "testing"

func TestPutFinalLoc(t *testing.T) {
	p := &put{loc: Loc{Off: 1}}
	h := &put{via: p}
	if p.finalLoc().Off != 1 || h.finalLoc().Off != 1 {
		t.Fatal("finalLoc")
	}
}
