//ff:func feature=delta type=decoder control=sequence
//ff:what TestDescribedNodes — LOCAL subjects grouped by their TRIPLEs in order; equal descriptions list both LOCALs; other records ignored
package delta

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDescribedNodes(t *testing.T) {
	body := []core.Record{
		tripleOf(1, 854, core.EncodeLocalRef(20), "u"),
		tripleOf(2, 854, core.EncodeLocalRef(21), "u"),
		tripleOf(3, 1476, core.EncodeLocalRef(20), "t"),
		tripleOf(4, 854, core.EncodeLocalRef(22), "u"),
		bindOf(5, 20, "x"),
	}
	n := describedNodes(body)
	_, u, _ := tripleAbout(body[0])
	_, ti, _ := tripleAbout(body[2])
	if len(n) != 2 || !slices.Equal(n[nodeKey([][]core.Word{u})], []uint64{21, 22}) || !slices.Equal(n[nodeKey([][]core.Word{u, ti})], []uint64{20}) {
		t.Fatalf("%v", n)
	}
}
