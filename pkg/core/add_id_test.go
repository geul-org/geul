//ff:func feature=core type=helper control=sequence
//ff:what TestAddID — a nil set is made; an existing set grows in place
package core

import "testing"

func TestAddID(t *testing.T) {
	s := addID(nil, 3)
	if !s[3] || len(s) != 1 {
		t.Fatalf("made: %v", s)
	}
	if addID(s, 5); !s[5] || len(s) != 2 {
		t.Fatalf("grown: %v", s)
	}
}
