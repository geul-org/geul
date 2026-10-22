//ff:func feature=inspect type=decoder control=sequence
//ff:what TestParseDocName — a 1-word Issuer splits; a 3-word Issuer fails
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseDocName(t *testing.T) {
	d, ok := parseDocName(core.Record{Payload: []core.Word{1, 0, 0, 0, 2}})
	if !ok || len(d.Issuer) != 1 || d.Issuer[0] != 1 || len(d.Doc) != 4 || d.Doc[3] != 2 {
		t.Fatalf("got %+v %v", d, ok)
	}
	if _, ok := parseDocName(core.Record{Payload: []core.Word{1, 1, 1, 0, 0, 0, 2}}); ok {
		t.Fatal("3-word issuer: want not ok")
	}
}
