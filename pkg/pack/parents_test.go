//ff:func feature=pack type=writer control=sequence
//ff:what TestParents — internal blocks of fanout children: first keys and written locations
package pack

import "testing"

func TestParents(t *testing.T) {
	bs := []block{{keys: []string{"a"}}, {keys: []string{"c"}}, {keys: []string{"e"}}}
	ps := parents(bs, []Loc{{Off: 1}, {Off: 2}, {Off: 3}}, blkHInner, 2)
	if len(ps) != 2 || ps[0].keys[1] != "c" || ps[0].locs[1].Off != 2 || ps[1].keys[0] != "e" || ps[1].kind != blkHInner {
		t.Fatal(ps)
	}
}
