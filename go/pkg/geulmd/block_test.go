//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestBlock — lines in a `geul` block; an empty block has no body line
package geulmd

import "testing"

func TestBlock(t *testing.T) {
	if got := block([]string{"A", "B"}); got != "```geul\nA\nB\n```\n" {
		t.Fatalf("got %q", got)
	}
	if got := block(nil); got != "```geul\n```\n" {
		t.Fatalf("empty: %q", got)
	}
}
