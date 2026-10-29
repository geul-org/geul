//ff:func feature=pack type=decoder control=sequence
//ff:what TestPKeyStep — inverts pKeyDelta; overflow and a zero slot step fail
package pack

import (
	"math"
	"testing"
)

func TestPKeyStep(t *testing.T) {
	if i, d, s, ok := pKeyStep(0, 0, 0, 0, 1, 5, 2); !ok || i != 1 || d != 5 || s != 2 {
		t.Fatal("first")
	}
	if i, _, _, ok := pKeyStep(1, 1, 5, 2, 2, 4, 1); !ok || i != 3 {
		t.Fatal("issuer")
	}
	if _, d, _, ok := pKeyStep(1, 1, 5, 2, 0, 1, 0); !ok || d != 6 {
		t.Fatal("doc")
	}
	if _, _, s, ok := pKeyStep(1, 1, 6, 0, 0, 0, 1); !ok || s != 1 {
		t.Fatal("slot")
	}
	if _, _, _, ok := pKeyStep(1, 1, 6, 3, 0, 0, 0); ok {
		t.Fatal("no increase")
	}
	if _, _, _, ok := pKeyStep(0, 0, 0, 0, math.MaxUint32+1, 0, 0); ok {
		t.Fatal("issuer overflow")
	}
	if _, _, _, ok := pKeyStep(1, 1, math.MaxUint64, 0, 0, 1, 0); ok {
		t.Fatal("doc overflow")
	}
}
