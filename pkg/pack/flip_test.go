//ff:func feature=pack type=helper control=sequence
//ff:what flip — invert one byte of a file in place
package pack

import (
	"os"
	"testing"
)

// flip XORs byte off of file with 0xFF.
func flip(t *testing.T, file string, off int) {
	t.Helper()
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	b[off] ^= 0xFF
	os.Chmod(file, 0o644)
	if err := os.WriteFile(file, b, 0o644); err != nil {
		t.Fatal(err)
	}
}
