//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteNew — a new file, an existing one refused, replaced with force
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteNew(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f")
	if err := writeNew(p, []byte("a"), false); err != nil {
		t.Fatal(err)
	}
	if err := writeNew(p, []byte("b"), false); err == nil {
		t.Fatal("exists: want error")
	}
	writeNew(p, []byte("c"), true)
	if b, _ := os.ReadFile(p); string(b) != "c" {
		t.Fatal(string(b))
	}
}
