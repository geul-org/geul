//ff:func feature=pack type=writer control=sequence
//ff:what TestCloseRuns — every run file is closed
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCloseRuns(t *testing.T) {
	f, err := os.Create(filepath.Join(t.TempDir(), "run"))
	if err != nil {
		t.Fatal(err)
	}
	closeRuns([]*runReader{{f: f}})
	if _, err := f.Write([]byte{1}); err == nil {
		t.Fatal("file still open")
	}
}
