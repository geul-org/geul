//ff:func feature=pack type=writer control=sequence
//ff:what TestSyncFile — nil is nothing to do, an open file syncs, a closed one fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSyncFile(t *testing.T) {
	if err := syncFile(nil); err != nil {
		t.Fatal(err)
	}
	f, _ := os.Create(filepath.Join(t.TempDir(), "f"))
	if err := syncFile(f); err != nil {
		t.Fatal(err)
	}
	f.Close()
	if err := syncFile(f); err == nil {
		t.Fatal("closed: want error")
	}
}
