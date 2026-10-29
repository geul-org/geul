//ff:func feature=pack type=helper control=sequence
//ff:what TestOsFSCreateTemp — a temp file in the directory; a missing directory fails
package pack

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestOsFSCreateTemp(t *testing.T) {
	d := t.TempDir()
	f, err := (osFS{}).CreateTemp(d)
	if err != nil || !strings.HasPrefix(filepath.Base(f.Name()), ".tmp-") {
		t.Fatal(err)
	}
	f.Close()
	if _, err := (osFS{}).CreateTemp(filepath.Join(d, "none")); err == nil {
		t.Fatal("missing dir")
	}
}
