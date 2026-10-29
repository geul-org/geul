//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what objPath — the local file of the first object of a kind in r's OBJECTS
package pack

import (
	"path/filepath"
	"testing"
)

// objPath finds the first object of kind in out's generation r.
func objPath(t *testing.T, r *Reader, out string, kind byte) string {
	t.Helper()
	for _, o := range r.m.objects {
		if o.kind == kind {
			return filepath.Join(out, filepath.FromSlash(objectName(o)))
		}
	}
	t.Fatal("no object of kind", kind)
	return ""
}
