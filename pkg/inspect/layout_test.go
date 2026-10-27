//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestLayout — every sample entry's parts rejoin to its words, head first
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLayout(t *testing.T) {
	d, err := Decode(sampleDoc(t))
	if err != nil {
		t.Fatal(err)
	}
	for i, e := range append(d.Head, d.Body...) {
		parts := Layout(e)
		var ws []core.Word
		for _, p := range parts {
			ws = append(ws, p.Words...)
		}
		if !slices.Equal(ws, e.Words) || parts[0].Name != "TAG" || !parts[0].Head {
			t.Fatalf("entry %d: %+v", i, parts)
		}
	}
}
