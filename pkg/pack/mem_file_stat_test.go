//ff:func feature=pack type=helper control=sequence
//ff:what TestMemFileStat — the stored info
package pack

import "testing"

func TestMemFileStat(t *testing.T) {
	fi, err := (&memFile{info: fileInfo{name: "a", size: 2}}).Stat()
	if err != nil || fi.Size() != 2 {
		t.Fatal(err)
	}
}
