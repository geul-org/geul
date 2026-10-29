//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoModTime — zero
package pack

import "testing"

func TestFileInfoModTime(t *testing.T) {
	if !(fileInfo{}).ModTime().IsZero() {
		t.Fatal("modTime")
	}
}
