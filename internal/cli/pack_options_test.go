//ff:func feature=cli type=helper control=sequence
//ff:what TestPackOptions — CLI writes verify each generation
package cli

import "testing"

func TestPackOptions(t *testing.T) {
	if !packOptions().Verify {
		t.Fatal("verify")
	}
}
