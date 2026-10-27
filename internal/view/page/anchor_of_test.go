//ff:func feature=page type=helper control=sequence
//ff:what TestAnchorOf — `e` plus the Edge Id
package page

import "testing"

func TestAnchorOf(t *testing.T) {
	if got := anchorOf(42); got != "e42" {
		t.Fatalf("got %q", got)
	}
}
