//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamTree — rootMax 0 or few heads give one root unwritten; more than rootMax are grouped by fanout into written blocks; one group left is the root; a failing codec fails
package pack

import "testing"

func TestStreamTree(t *testing.T) {
	heads := []streamKV{{Key: "a"}, {Key: "b"}, {Key: "c"}, {Key: "d"}, {Key: "e"}}
	s := streamTestNew(t, nil)
	if b, err := s.tree(heads, blkHInner, 2, 0); err != nil || len(b.keys) != 5 || s.st.Index.Num != -1 {
		t.Fatal("rootMax 0")
	}
	b, err := s.tree(heads, blkHInner, 2, 2)
	if err != nil || len(b.keys) != 2 || b.keys[1] != "e" || s.st.Index.Count != 5 {
		t.Fatalf("%v %d %d", err, len(b.keys), s.st.Index.Count)
	}
	one := streamTestNew(t, nil)
	if b, err := one.tree(heads, blkHInner, 8, 2); err != nil || len(b.keys) != 5 || one.st.Index.Num != -1 {
		t.Fatal("one group is the root")
	}
	if _, err := streamTestNew(t, failCodec{}).tree(heads, blkHInner, 2, 2); err == nil {
		t.Fatal("failing codec: want error")
	}
}
