//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestDeclaredFields — fields in order; a missing FIELD fails
package codecspec

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDeclaredFields(t *testing.T) {
	fields := map[uint64]*Field{1: {Name: "a"}, 2: {Name: "b"}}
	if fs, ok := declaredFields(core.Declare{Fields: []uint64{2, 1}}, fields); !ok || fs[0].Name != "b" || fs[1].Name != "a" {
		t.Fatalf("got %+v %v", fs, ok)
	}
	if _, ok := declaredFields(core.Declare{Fields: []uint64{3}}, fields); ok {
		t.Fatal("missing: want false")
	}
}
