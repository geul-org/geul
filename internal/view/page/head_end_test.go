//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestHeadEnd — the second node (object), else the only node (subject), else none; literals never
package page

import "testing"

func TestHeadEnd(t *testing.T) {
	res, lit := GraphNode{Class: "res"}, GraphNode{Class: "lit"}
	for want, ends := range map[int][]GraphNode{-1: {lit}, 1: {lit, res}, 2: {res, lit, res, res}} {
		if got := headEnd(ends); got != want {
			t.Fatalf("%v: got %d want %d", ends, got, want)
		}
	}
}
