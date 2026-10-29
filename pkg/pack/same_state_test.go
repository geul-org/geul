//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what TestSameState — equal maps pass; a different size, a missing path or a different location fail
package pack

import (
	"errors"
	"testing"
)

func TestSameState(t *testing.T) {
	a := map[string]Loc{"x": {Off: 1}}
	if sameState(a, map[string]Loc{"x": {Off: 1}}) != nil {
		t.Fatal("equal")
	}
	for _, b := range []map[string]Loc{{}, {"y": {Off: 1}}, {"x": {Off: 2}}} {
		if !errors.Is(sameState(a, b), ErrCorrupt) {
			t.Fatal(b)
		}
	}
}
