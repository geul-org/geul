//ff:func feature=resolve type=resolver control=sequence
//ff:what TestRetiredIn — a kept ID and an ID from NEXT-LOCAL-ID on pass; a dropped ID, a skipped one and any ID below next with no previous version are found
package resolve

import "testing"

func TestRetiredIn(t *testing.T) {
	prev := edgeDoc(7, "a", "b")
	if _, ok := retiredIn(prev, edgeDoc(7, "a", "b", "c"), 3); ok {
		t.Fatal("#3 from NEXT-LOCAL-ID: want none")
	}
	if id, ok := retiredIn(edgeDoc(7, "a"), edgeDoc(7, "a", "b"), 3); !ok || id != 2 {
		t.Fatalf("#2 dropped: %d %v", id, ok)
	}
	if id, ok := retiredIn(nil, edgeDoc(7, "a"), 2); !ok || id != 1 {
		t.Fatalf("no previous version: %d %v", id, ok)
	}
}
