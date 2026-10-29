//ff:func feature=pack type=helper control=sequence
//ff:what TestFileInfoIsDir — the directory flag
package pack

import "testing"

func TestFileInfoIsDir(t *testing.T) {
	if !(fileInfo{dir: true}).IsDir() || (fileInfo{}).IsDir() {
		t.Fatal("isDir")
	}
}
