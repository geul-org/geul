//ff:func feature=pack type=helper control=sequence
//ff:what TestDirEntryIsDir — the directory flag
package pack

import "testing"

func TestDirEntryIsDir(t *testing.T) {
	if !(dirEntry{dir: true}).IsDir() || (dirEntry{}).IsDir() {
		t.Fatal("isDir")
	}
}
