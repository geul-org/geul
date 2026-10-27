//ff:func feature=config type=helper control=sequence
//ff:what TestDiscover — the nearest folder holding .well-known/geul from a nested folder; none in a plain tree
package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscover(t *testing.T) {
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, ".well-known", "geul"), 0o755)
	deep := filepath.Join(root, "a", "b")
	os.MkdirAll(deep, 0o755)
	if got := Discover(deep); got != root {
		t.Fatalf("got %q", got)
	}
	if got := Discover(t.TempDir()); got != "" && got == root {
		t.Fatalf("plain tree: %q", got)
	}
}
