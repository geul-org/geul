//ff:func feature=core type=encoder control=sequence
//ff:what TestPrependTag — tag then body; a body over the record limit is refused
package core

import "testing"

func TestPrependTag(t *testing.T) {
	ws, err := prependTag(KindEdge, 1, 2, []Word{7, 8})
	if err != nil || len(ws) != 3 || ws[1] != 7 {
		t.Fatal(ws, err)
	}
	if _, err := prependTag(KindEdge, 1, 1<<40, nil); err == nil {
		t.Fatal("over 2^32-1 words: want error")
	}
}
