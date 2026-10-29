//ff:func feature=pack type=helper control=sequence
//ff:what TestHasPrefix — any sorted path under the prefix
package pack

import "testing"

func TestHasPrefix(t *testing.T) {
	ps := []string{"a/b", "c"}
	if !hasPrefix(ps, "a/") || hasPrefix(ps, "b/") || hasPrefix(ps, "d") {
		t.Fatal("hasPrefix")
	}
}
