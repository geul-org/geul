//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestIsSlug — slugs pass; empty, dashes at the ends, 0x prefix, uppercase and other characters fail
package schema

import (
	"testing"
)

func TestIsSlug(t *testing.T) {
	for s, want := range map[string]bool{"this-is-sample-doc": true, "a1": true, "": false, "-a": false, "a-": false, "0x12": false, "Abc": false, "a_b": false, "민법": false} {
		if isSlug(s) != want {
			t.Fatalf("%q: want %v", s, want)
		}
	}
}
