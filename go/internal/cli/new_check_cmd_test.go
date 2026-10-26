//ff:func feature=cli type=command control=sequence
//ff:what TestNewCheckCmd — a valid archive passes --archive; a broken file is a finding and fails
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewCheckCmd(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("check", "--archive", draft)
	bad := filepath.Join(h.data, "bad.geul")
	os.WriteFile(bad, []byte{0, 5}, 0o644)
	out, err := h.run("", "check", bad)
	if err == nil || !strings.HasPrefix(out, bad+": ") {
		t.Fatalf("%q %v", out, err)
	}
}
