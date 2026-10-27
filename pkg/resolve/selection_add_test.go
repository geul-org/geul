//ff:func feature=resolve type=helper control=sequence
//ff:what TestSelectionAdd — a new version is kept; a latest request landing on a fetched version reuses that Node
package resolve

import (
	"testing"
)

func TestSelectionAdd(t *testing.T) {
	s := selection{nodes: map[NodeKey]*Node{}, latest: map[DocKey]*Node{}}
	pinned := &Node{Key: sampleKey(1, 2), Version: 2}
	if got := s.add(sampleKey(1, 2), pinned); got != pinned || len(s.latest) != 0 {
		t.Fatalf("pinned: %p", got)
	}
	again := &Node{Key: sampleKey(1, 0), Version: 2}
	if got := s.add(sampleKey(1, 0), again); got != pinned || s.latest[sampleKey(1, 0).DocKey()] != pinned || len(s.nodes) != 1 {
		t.Fatalf("latest reuse: %p", got)
	}
}
