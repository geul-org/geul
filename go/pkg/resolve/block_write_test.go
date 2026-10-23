//ff:func feature=resolve type=resolver control=selection
//ff:what blockWrite — test fixture: make a later write to a Dir path fail without breaking reads before it
package resolve

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// blockWrite makes writing path in d fail. A document path becomes a directory.
// A version list path becomes a dangling symlink, so reading the list still
// finds none but writing it fails.
func blockWrite(t *testing.T, d Dir, path string) {
	t.Helper()
	f := d.file(path)
	if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
		t.Fatal(err)
	}
	var err error
	switch {
	case strings.HasSuffix(path, "/"):
		err = os.Symlink(filepath.Join(d.Root, "no", "such"), f)
	default:
		err = os.MkdirAll(f, 0o755)
	}
	if err != nil {
		t.Fatal(err)
	}
}
