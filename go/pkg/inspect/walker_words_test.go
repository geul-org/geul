//ff:func feature=inspect type=decoder control=sequence
//ff:what TestWalkerWords — n words become one part; too few words or a bad walk adds nothing
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestWalkerWords(t *testing.T) {
	w := &walker{ws: []core.Word{1, 2, 3}}
	w.words("a", 2)
	if len(w.out) != 1 || len(w.out[0].Words) != 2 || w.out[0].Record || w.pos != 2 {
		t.Fatalf("got %+v", w)
	}
	w.words("b", 2)
	if !w.bad || len(w.out) != 1 {
		t.Fatalf("too few: %+v", w)
	}
}
