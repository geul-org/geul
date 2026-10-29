//ff:func feature=cli type=helper control=sequence
//ff:what TestDictNote — empty without a new dictionary, its ID otherwise
package cli

import "testing"

func TestDictNote(t *testing.T) {
	if dictNote(0) != "" || dictNote(32768) != ", dictionary 32768" {
		t.Fatal("dictNote")
	}
}
