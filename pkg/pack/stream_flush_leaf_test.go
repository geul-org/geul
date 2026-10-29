//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamFlushLeaf — the leaf is written, its head kept, the pending leaf emptied; a failing codec fails
package pack

import "testing"

func TestStreamFlushLeaf(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.PLeaf = []streamKV{{Key: pKey(0, 1, 0)}, {Key: pKey(0, 1, 1)}}
	if err := s.flushLeaf(); err != nil || len(s.st.PHeads) != 1 || s.st.PHeads[0].Key != pKey(0, 1, 0) || len(s.st.PLeaf) != 0 {
		t.Fatalf("%v", err)
	}
	f := streamTestNew(t, failCodec{})
	f.st.PLeaf = []streamKV{{Key: pKey(0, 1, 0)}}
	if err := f.flushLeaf(); err == nil {
		t.Fatal("failing codec: want error")
	}
}
