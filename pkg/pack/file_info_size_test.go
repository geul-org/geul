//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoSize — the size
package pack

import "testing"

func TestFileInfoSize(t *testing.T) {
	if (fileInfo{size: 3}).Size() != 3 {
		t.Fatal("size")
	}
}
