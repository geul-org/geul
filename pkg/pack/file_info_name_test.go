//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoName — the name
package pack

import "testing"

func TestFileInfoName(t *testing.T) {
	if (fileInfo{name: "a"}).Name() != "a" {
		t.Fatal("name")
	}
}
