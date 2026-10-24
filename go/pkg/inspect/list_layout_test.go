//ff:func feature=inspect type=decoder control=sequence
//ff:what TestListLayout — the count Value then one REF part per item
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestListLayout(t *testing.T) {
	w := &walker{ws: []core.Word{0x0882, 2, 0x0088, 5, 0x0088, 6}}
	listLayout(w, "with", 2)
	if w.bad || w.pos != 6 || len(w.out) != 3 || w.out[0].Name != "with.count" || w.out[2].Name != "with" {
		t.Fatalf("got %+v", w)
	}
}
