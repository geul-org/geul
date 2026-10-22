//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderDocName — issuer and doc id render; an impossible issuer width does not
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestRenderDocName(t *testing.T) {
	if got, ok := renderDocName(core.Record{Payload: []core.Word{0, 1, 0, 0, 0, 9}}); !ok || got != "DOC-NAME 0x00000001 0x0000000000000009" {
		t.Fatalf("got %q %v", got, ok)
	}
	if _, ok := renderDocName(core.Record{Payload: []core.Word{0, 0, 0, 0, 0, 0, 0}}); ok {
		t.Fatal("3-word issuer: want not ok")
	}
}
