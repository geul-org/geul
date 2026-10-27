//ff:func feature=delta type=helper control=sequence
//ff:what TestEdgeKey — same codec and words give the same key; another codec or word differs
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeKey(t *testing.T) {
	a, b := edgeKey(3, []core.Word{1, 2}), edgeKey(3, []core.Word{1, 2})
	if a != b || a == edgeKey(4, []core.Word{1, 2}) || a == edgeKey(3, []core.Word{1, 3}) {
		t.Fatal("keys")
	}
}
