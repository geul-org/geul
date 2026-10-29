//ff:func feature=pack type=helper control=sequence
//ff:what fixtureFile — the bytes of one testdata archive path
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

// fixtureFile reads path (relative to `.well-known/geul/`) from testdata.
func fixtureFile(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata/archive/.well-known/geul", filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return b
}
