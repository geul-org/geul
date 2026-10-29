//ff:func feature=pack type=helper control=sequence
//ff:what TestDirFileStat — the directory info
package pack

import "testing"

func TestDirFileStat(t *testing.T) {
	fi, err := (&dirFile{info: fileInfo{name: "d", dir: true}}).Stat()
	if err != nil || !fi.IsDir() {
		t.Fatal(err)
	}
}
