//ff:func feature=resolve type=helper control=sequence
//ff:what TestKeepHighest — first sets, higher replaces, lower keeps
package resolve

import (
	"testing"
)

func TestKeepHighest(t *testing.T) {
	sel := map[DocKey]*Node{}
	keepHighest(sel, &Node{Key: sampleKey(1, 2), Version: 2})
	keepHighest(sel, &Node{Key: sampleKey(1, 1), Version: 1})
	if sel[sampleKey(1, 0).DocKey()].Version != 2 {
		t.Fatal("lower replaced higher")
	}
	keepHighest(sel, &Node{Key: sampleKey(1, 3), Version: 3})
	if sel[sampleKey(1, 0).DocKey()].Version != 3 {
		t.Fatal("higher not kept")
	}
}
