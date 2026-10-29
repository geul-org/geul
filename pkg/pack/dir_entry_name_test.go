//ff:func feature=pack type=helper control=sequence
//ff:what TestDirEntryName — the name
package pack

import "testing"

func TestDirEntryName(t *testing.T) {
	if (dirEntry{name: "a"}).Name() != "a" {
		t.Fatal("name")
	}
}
