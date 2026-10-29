//ff:func feature=pack type=helper control=sequence
//ff:what TestOsFSRename — moves the file readable to all; a missing source fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOsFSRename(t *testing.T) {
	d := t.TempDir()
	a, b := filepath.Join(d, "a"), filepath.Join(d, "b")
	os.WriteFile(a, []byte("x"), 0o600)
	if err := (osFS{}).Rename(a, b); err != nil {
		t.Fatal(err)
	}
	if st, err := os.Stat(b); err != nil || st.Mode().Perm() != 0o644 {
		t.Fatal(err)
	}
	if err := (osFS{}).Rename(a, b); err == nil {
		t.Fatal("missing")
	}
}
