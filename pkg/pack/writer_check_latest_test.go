//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterCheckLatest — passes while latest.txt names the starting generation (or is absent for the first); ErrConflict otherwise
package pack

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriterCheckLatest(t *testing.T) {
	if err := (&writer{out: t.TempDir()}).checkLatest(); err != nil {
		t.Fatal(err)
	}
	_, out := built(t, small())
	w := testWriter(t, out, Batch{})
	if err := w.checkLatest(); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(out, latestFile), []byte("2 x\n"), 0o644)
	if err := w.checkLatest(); !errors.Is(err, ErrConflict) {
		t.Fatal(err)
	}
}
