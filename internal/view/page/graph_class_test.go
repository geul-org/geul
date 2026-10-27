//ff:func feature=page type=helper control=sequence
//ff:what TestGraphClass — known node and edge names map to their class; unknown names are "other"
package page

import "testing"

func TestGraphClass(t *testing.T) {
	if graphClass("COUNTER-DECLARE") != "counter" || graphClass("DO-ONCE") != "run" || graphClass("SAME-AS") != "def" {
		t.Fatal("class mismatch")
	}
	if got := graphClass("LABEL"); got != "other" {
		t.Fatalf("unknown: %q", got)
	}
}
