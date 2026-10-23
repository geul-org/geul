//ff:func feature=inspect type=decoder control=sequence
//ff:what TestWalkerSpan — a SPAN is two parts at the tried width
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestWalkerSpan(t *testing.T) {
	w := &walker{ws: []core.Word{0, 1, 0, 2}, sw: 2}
	w.span("span")
	if len(w.out) != 2 || w.out[0].Name != "span.off" || w.out[1].Name != "span.len" || w.pos != 4 {
		t.Fatalf("got %+v", w)
	}
}
