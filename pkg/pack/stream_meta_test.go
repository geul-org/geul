//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamMeta — nil before a checkpoint, then the last checkpoint's meta
package pack

import "testing"

func TestStreamMeta(t *testing.T) {
	s := streamTestNew(t, nil)
	if s.Meta() != nil {
		t.Fatal("before")
	}
	s.Checkpoint([]byte("p"))
	if string(s.Meta()) != "p" {
		t.Fatal(string(s.Meta()))
	}
}
