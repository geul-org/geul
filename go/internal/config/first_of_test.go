//ff:func feature=config type=helper control=sequence
//ff:what TestFirstOf — the first non-empty; none is empty
package config

import "testing"

func TestFirstOf(t *testing.T) {
	if firstOf("", "a", "b") != "a" || firstOf("", "") != "" {
		t.Fatal("firstOf")
	}
}
