//ff:func feature=cli type=command control=sequence
//ff:what TestNewHashCmd — hash of a version equals the hash of its file
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewHashCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	f := filepath.Join(h.data, "v2.geul")
	os.WriteFile(f, []byte(h.must("cat", addr+"@2")), 0o644)
	if a, b := h.must("hash", addr+"@2"), h.must("hash", f); a != b || len(a) != 67 {
		t.Fatalf("%q %q", a, b)
	}
}
