//ff:func feature=pack type=helper control=sequence
//ff:what TestSumFile — SHA-256 of a prefix; a read failure is returned
package pack

import (
	"crypto/sha256"
	"testing"
)

func TestSumFile(t *testing.T) {
	f, _ := osFS{}.CreateTemp(t.TempDir())
	f.Write([]byte("abcdef"))
	if sum, err := sumFile(f, 3); err != nil || sum != sha256.Sum256([]byte("abc")) {
		t.Fatal(err)
	}
	if _, err := sumFile(&badFile{objFile: f, readAt: true}, 3); err == nil {
		t.Fatal("read error")
	}
}
