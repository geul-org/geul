//ff:func feature=schema type=helper control=sequence
//ff:what TestWithID — a nil set is made; an existing set grows in place
package schema

import "testing"

func TestWithID(t *testing.T) {
	s := withID(nil, 3)
	if !s[3] || len(s) != 1 {
		t.Fatalf("made: %v", s)
	}
	if withID(s, 5); !s[5] || len(s) != 2 {
		t.Fatalf("grown: %v", s)
	}
}
