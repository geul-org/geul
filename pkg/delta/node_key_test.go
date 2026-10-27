//ff:func feature=delta type=helper control=sequence
//ff:what TestNodeKey — equal descriptions share a key; order and content tell them apart
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNodeKey(t *testing.T) {
	a := [][]core.Word{{1, 2}, {3}}
	if nodeKey(a) != nodeKey([][]core.Word{{1, 2}, {3}}) || nodeKey(a) == nodeKey([][]core.Word{{3}, {1, 2}}) || nodeKey(a) == nodeKey([][]core.Word{{1, 2, 3}}) {
		t.Fatal(nodeKey(a))
	}
}
