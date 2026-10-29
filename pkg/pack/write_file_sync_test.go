//ff:func feature=pack type=writer control=sequence
//ff:what TestWriteFileSync — the file appears with its bytes, readable by all; a missing folder and a target that is a folder fail without leaving a temp file
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileSync(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "latest.txt")
	if err := writeFileSync(p, []byte("1 x\n")); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(p); string(b) != "1 x\n" {
		t.Fatal(string(b))
	}
	if fi, _ := os.Stat(p); fi.Mode().Perm() != 0o644 {
		t.Fatal(fi.Mode())
	}
	if err := writeFileSync(filepath.Join(dir, "no", "f"), nil); err == nil {
		t.Fatal("missing folder: want error")
	}
	os.MkdirAll(filepath.Join(dir, "d", "x"), 0o755)
	if err := writeFileSync(filepath.Join(dir, "d"), nil); err == nil {
		t.Fatal("folder target: want error")
	}
	if ents, _ := os.ReadDir(dir); len(ents) != 2 {
		t.Fatalf("temp file left: %d", len(ents))
	}
}
