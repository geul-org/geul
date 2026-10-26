//ff:func feature=cli type=helper control=sequence
//ff:what TestEdgeIndex — the Edge with the ID; non-Edges and missing IDs are skipped
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/notation"
)

func TestEdgeIndex(t *testing.T) {
	body := []core.Record{core.EncodeSpecVersion(2), edgeRecord(3, notation.Edge{Codec: 4}), edgeRecord(5, notation.Edge{Codec: 4})}
	if edgeIndex(body, 5) != 2 || edgeIndex(body, 9) != -1 {
		t.Fatal("index")
	}
}
