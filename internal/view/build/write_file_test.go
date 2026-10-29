//ff:func feature=build type=helper control=sequence
//ff:what TestWriteFile — the writer fills the file; an uncreatable path, a failed flush and a writer error fail
package build

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f")
	if err := writeFile(p, func(w io.Writer) error { _, err := io.WriteString(w, "ok"); return err }); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(p); string(b) != "ok" {
		t.Fatalf("got %q", b)
	}
	if err := writeFile(filepath.Join(p, "x"), func(io.Writer) error { return nil }); err == nil {
		t.Fatal("under a file: want error")
	}
	if err := writeFile("/dev/full", func(w io.Writer) error { _, err := io.WriteString(w, "ok"); return err }); err == nil {
		t.Fatal("flush to a full device: want error")
	}
	bad := errors.New("bad")
	if err := writeFile(p, func(io.Writer) error { return bad }); err != bad {
		t.Fatalf("writer: %v", err)
	}
}
