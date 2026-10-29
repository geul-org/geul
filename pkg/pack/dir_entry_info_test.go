//ff:func feature=pack type=helper control=sequence
//ff:what TestDirEntryInfo — the same name and kind as a FileInfo
package pack

import "testing"

func TestDirEntryInfo(t *testing.T) {
	fi, err := (dirEntry{name: "a", dir: true}).Info()
	if err != nil || fi.Name() != "a" || !fi.IsDir() {
		t.Fatal(err)
	}
}
