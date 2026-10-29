//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHRoot — no key no layer; one leaf is the root; written leaves get a tree; a failing codec and a failing tree fail
package pack

import "testing"

func TestStreamHRoot(t *testing.T) {
	s := streamTestNew(t, nil)
	if ls, err := s.hRoot(nil, nil, 0); ls != nil || err != nil {
		t.Fatal("no key")
	}
	leaf := []streamKV{{Key: "k"}}
	if ls, err := s.hRoot(nil, leaf, 1); err != nil || ls[0].count != 1 || s.st.Index.Num != -1 {
		t.Fatal("single leaf is the root")
	}
	heads := []streamKV{{Key: "a"}, {Key: "b"}}
	if ls, err := s.hRoot(heads, leaf, 9); err != nil || ls[0].count != 9 || s.st.Index.Count != 1 {
		t.Fatalf("%v", err)
	}
	if _, err := streamTestNew(t, failCodec{}).hRoot(heads, leaf, 9); err == nil {
		t.Fatal("failing codec: want error")
	}
	tr2 := streamTestNew(t, failCodecAfter(1))
	tr2.o.Build.HRoot, tr2.o.Build.HFanout = 1, 2
	if _, err := tr2.hRoot(heads, leaf, 9); err == nil {
		t.Fatal("tree failure: want error")
	}
}
