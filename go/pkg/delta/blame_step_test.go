//ff:func feature=delta type=decoder control=sequence
//ff:what TestBlameStep — new IDs are added at v, equal content keeps its versions, different content is changed at v
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestBlameStep(t *testing.T) {
	seen := map[uint64]*Line{}
	s, _ := splitSnapshot(encodeDoc(nil, []core.Record{bindOf(1, 9, "a")}))
	blameStep(seen, s, 1)
	blameStep(seen, s, 2)
	if seen[1].Added != 1 || seen[1].Changed != 1 {
		t.Fatalf("same: %+v", seen[1])
	}
	s2, _ := splitSnapshot(encodeDoc(nil, []core.Record{bindOf(1, 9, "x")}))
	blameStep(seen, s2, 3)
	if seen[1].Changed != 3 || seen[1].Added != 1 {
		t.Fatalf("changed: %+v", seen[1])
	}
}
