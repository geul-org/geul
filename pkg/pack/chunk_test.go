//ff:func feature=pack type=writer control=sequence
//ff:what TestChunk — leaves of at most size entries
package pack

import "testing"

func TestChunk(t *testing.T) {
	bs := chunk([]kv{{"a", Loc{}}, {"b", Loc{}}, {"c", Loc{}}}, blkHLeaf, 2)
	if len(bs) != 2 || len(bs[0].keys) != 2 || bs[1].keys[0] != "c" || bs[1].kind != blkHLeaf {
		t.Fatal(bs)
	}
}
