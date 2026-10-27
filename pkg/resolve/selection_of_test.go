//ff:func feature=resolve type=helper control=sequence
//ff:what TestSelectionOf — pinned Keys by exact version, latest Keys by the latest binding, nil when unbound
package resolve

import (
	"testing"
)

func TestSelectionOf(t *testing.T) {
	a1, a2 := &Node{Key: sampleKey(1, 1), Version: 1}, &Node{Key: sampleKey(1, 0), Version: 2}
	d := sampleKey(1, 0).DocKey()
	s := selection{nodes: map[NodeKey]*Node{{DocKey: d, Version: 1}: a1, {DocKey: d, Version: 2}: a2}, latest: map[DocKey]*Node{d: a2}}
	if s.of(sampleKey(1, 1)) != a1 || s.of(sampleKey(1, 0)) != a2 || s.of(sampleKey(1, 3)) != nil || s.of(sampleKey(2, 0)) != nil {
		t.Fatal("wrong binding")
	}
}
