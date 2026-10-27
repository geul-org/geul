//ff:func feature=page type=helper control=sequence
//ff:what TestAddNodes — new ids appended in order, seen ids skipped
package page

import "testing"

func TestAddNodes(t *testing.T) {
	seen := map[string]bool{"a": true}
	got := addNodes(nil, seen, []GraphNode{{ID: "a"}, {ID: "b"}, {ID: "b"}})
	if len(got) != 1 || got[0].ID != "b" || !seen["b"] {
		t.Fatalf("got %+v", got)
	}
}
