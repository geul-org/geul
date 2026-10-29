//ff:func feature=pack type=helper control=sequence
//ff:what TestOsFSRemove — removes a file
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOsFSRemove(t *testing.T) {
	a := filepath.Join(t.TempDir(), "a")
	os.WriteFile(a, nil, 0o644)
	if err := (osFS{}).Remove(a); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(a); err == nil {
		t.Fatal("still there")
	}
}
