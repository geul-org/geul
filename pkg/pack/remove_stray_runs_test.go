//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestRemoveStrayRuns — runs the checkpoint lists and other files stay, others go; an unreadable folder and an unremovable stray fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveStrayRuns(t *testing.T) {
	out := t.TempDir()
	dir := filepath.Join(out, streamDir)
	os.MkdirAll(dir, 0o755)
	for _, n := range []string{"run-000000", "run-000001", "state"} {
		os.WriteFile(filepath.Join(dir, n), nil, 0o644)
	}
	if err := removeStrayRuns(out, []string{"run-000000"}); err != nil {
		t.Fatal(err)
	}
	for n, want := range map[string]bool{"run-000000": true, "run-000001": false, "state": true} {
		if _, err := os.Stat(filepath.Join(dir, n)); (err == nil) != want {
			t.Errorf("%s: kept %v", n, err == nil)
		}
	}
	if err := removeStrayRuns(t.TempDir(), nil); err == nil {
		t.Fatal("no working folder: want error")
	}
	os.MkdirAll(filepath.Join(dir, "run-000002", "x"), 0o755)
	if err := removeStrayRuns(out, nil); err == nil {
		t.Fatal("unremovable stray: want error")
	}
}
