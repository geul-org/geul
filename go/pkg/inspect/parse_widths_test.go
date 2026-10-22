//ff:func feature=inspect type=decoder control=sequence
//ff:what TestParseWidths — the three width codes; a reserved bit or a wrong length fails
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseWidths(t *testing.T) {
	if w, ok := parseWidths(core.Record{Payload: []core.Word{0x6400}}); !ok || w != (Widths{Edge: 2, Var: 4, Import: 2}) {
		t.Fatalf("got %+v %v", w, ok)
	}
	if _, ok := parseWidths(core.Record{Payload: []core.Word{0x0001}}); ok {
		t.Fatal("reserved bit: want not ok")
	}
	if _, ok := parseWidths(core.Record{}); ok {
		t.Fatal("empty: want not ok")
	}
}
