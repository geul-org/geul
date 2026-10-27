//ff:func feature=cli type=command control=sequence
//ff:what TestNewBuildCmd — build needs --out, defaults to the data folder, takes a given tree, and fails without a data folder
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewBuildCmd(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "build"); err == nil || !strings.Contains(err.Error(), "out") {
		t.Fatalf("no --out: %v", err)
	}
	dst := t.TempDir()
	if out := h.must("build", "-o", dst); !strings.Contains(out, "0 documents, 1 directories") {
		t.Fatalf("data folder: %q", out)
	}
	if _, err := os.Stat(filepath.Join(dst, ".well-known", "geul", "index.html")); err != nil {
		t.Fatal(err)
	}
	if out := h.must("build", "-o", t.TempDir(), h.data); !strings.Contains(out, "1 directories") {
		t.Fatalf("given: %q", out)
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "build", "-o", t.TempDir()); err == nil || !strings.Contains(err.Error(), "no data folder") {
		t.Fatalf("no data folder: %v", err)
	}
}
