//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderLines — one line per entry, in order
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/inspect"
)

func TestRenderLines(t *testing.T) {
	got := renderLines([]inspect.Entry{{Name: "a"}, {Name: "b"}}, func(e inspect.Entry) string { return e.Name })
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("got %q", got)
	}
}
