//ff:func feature=resolve type=resolver control=sequence
//ff:what TestNamesFor — every name of one ID, sorted
package resolve

import (
	"slices"
	"testing"
)

func TestNamesFor(t *testing.T) {
	got := namesFor(map[string]string{"b": "x", "a": "x", "c": "y"}, "x")
	if !slices.Equal(got, []string{"a", "b"}) {
		t.Fatal(got)
	}
}
