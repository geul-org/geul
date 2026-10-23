//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestPublicationLayout — each version list record splits into ID, doc, version and its tail
package inspect

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestPublicationLayout(t *testing.T) {
	doc := []core.Word{0, 0, 0, 1}
	for want, p := range map[string]core.Publication{
		"ID doc version hash":   {Codec: core.EdgePublication, EdgeID: 1, Doc: doc, Version: 1},
		"ID doc version":        {Codec: core.EdgeDeprecation, EdgeID: 2, Doc: doc, Version: 1},
		"ID doc version reason": {Codec: core.EdgeTombstone, EdgeID: 3, Doc: doc, Version: 1, Reason: 2},
	} {
		w := &walker{ws: core.EncodePublication(p).Payload}
		publicationLayout(w, p)
		var names []string
		for _, pt := range w.out {
			names = append(names, pt.Name)
		}
		if got := strings.Join(names, " "); got != want || w.bad || w.pos != len(w.ws) {
			t.Fatalf("got %q bad %v", got, w.bad)
		}
	}
}
