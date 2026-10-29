//ff:func feature=pack type=helper control=sequence
//ff:what TestOsFSMkdirAll — nested directories are created
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOsFSMkdirAll(t *testing.T) {
	d := filepath.Join(t.TempDir(), "a", "b")
	if err := (osFS{}).MkdirAll(d); err != nil {
		t.Fatal(err)
	}
	if st, err := os.Stat(d); err != nil || !st.IsDir() {
		t.Fatal(err)
	}
}
