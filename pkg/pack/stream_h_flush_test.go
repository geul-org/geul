//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHFlush — a leaf short of HLeaf stays; a full one is written and its head added; a failing codec fails
package pack

import "testing"

func TestStreamHFlush(t *testing.T) {
	s := streamTestNew(t, nil)
	leaf := []streamKV{{Key: "a"}, {Key: "b"}}
	h, l, err := s.hFlush(nil, leaf)
	if err != nil || h != nil || len(l) != 2 {
		t.Fatal("short leaf")
	}
	leaf = append(leaf, streamKV{Key: "c"}, streamKV{Key: "d"})
	h, l, err = s.hFlush(nil, leaf)
	if err != nil || len(h) != 1 || h[0].Key != "a" || len(l) != 0 {
		t.Fatalf("%v", err)
	}
	if _, _, err := streamTestNew(t, failCodec{}).hFlush(nil, leaf); err == nil {
		t.Fatal("failing codec: want error")
	}
}
