//ff:func feature=cli type=command control=sequence
//ff:what TestAppCheck — a sound archive passes; nothing to check fails; a broken file in the archive fails; findings as JSON; a missing path and an archive without data folder fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppCheck(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if _, err := h.run("", "check", "--archive"); err != nil {
		t.Fatal(err)
	}
	if _, err := h.run("", "check"); err == nil {
		t.Fatal("nothing: want error")
	}
	bad := filepath.Join(h.data, "bad.geul")
	os.WriteFile(bad, []byte{0, 5}, 0o644)
	if out, _ := h.run("", "check", bad, "--format", "json"); !strings.Contains(out, `"path"`) {
		t.Fatal(out)
	}
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "bad.geul"), []byte{0, 5}, 0o644)
	if _, err := h.run("", "check", "--archive"); err == nil {
		t.Fatal("broken file in the archive: want error")
	}
	if _, err := h.run("", "check", filepath.Join(h.data, "none")); err == nil {
		t.Fatal("missing path: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "check", "--archive"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
