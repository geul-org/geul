//ff:func feature=pack type=helper control=sequence
//ff:what sumFile — SHA-256 of the first n bytes of a file
package pack

import (
	"crypto/sha256"
	"io"
)

// sumFile hashes f[0:n].
func sumFile(f io.ReaderAt, n uint64) ([32]byte, error) {
	h := sha256.New()
	var sum [32]byte
	if _, err := io.Copy(h, io.NewSectionReader(f, 0, int64(n))); err != nil {
		return sum, err
	}
	copy(sum[:], h.Sum(nil))
	return sum, nil
}
