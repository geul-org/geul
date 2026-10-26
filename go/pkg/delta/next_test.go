//ff:func feature=delta type=helper control=sequence
//ff:what TestNext — the next free ID, at least 1
package delta

import "testing"

func TestNext(t *testing.T) {
	if l, _ := NewLocalIDs(nil, 0); l.Next() != 1 {
		t.Fatal("want 1")
	}
}
