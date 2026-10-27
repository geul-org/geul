//ff:func feature=build type=helper control=sequence
//ff:what TestWrite — the vendor scripts are copied into a new directory; a directory that cannot be made, or a file that cannot be written, fails
package static

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "a", "_static")
	if err := Write(dir); err != nil {
		t.Fatal(err)
	}
	if b, err := os.ReadFile(filepath.Join(dir, "dagre.min.js")); err != nil || len(b) == 0 {
		t.Fatalf("dagre: %v", err)
	}
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	if err := Write(filepath.Join(file, "x")); err == nil {
		t.Fatal("under a file: want error")
	}
	blocked := t.TempDir()
	os.Mkdir(filepath.Join(blocked, "dagre.min.js"), 0o755)
	if err := Write(blocked); err == nil {
		t.Fatal("blocked file: want error")
	}
}
