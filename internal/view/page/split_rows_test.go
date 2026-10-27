//ff:func feature=page type=builder control=sequence
//ff:what TestSplitRows — NODE rows and Edge rows go to their group in order; other rows are dropped
package page

import "testing"

func TestSplitRows(t *testing.T) {
	n, e := splitRows([]Row{{Class: "node", Index: 0}, {Class: "link", Index: 1, Edge: true}, {Class: "value", Index: 2}, {Class: "node", Index: 3}, {Class: "raw", Index: 4, Edge: true}})
	if len(n) != 2 || n[1].Index != 3 || len(e) != 2 || e[0].Index != 1 || e[1].Index != 4 {
		t.Fatalf("got %v %v", n, e)
	}
}
