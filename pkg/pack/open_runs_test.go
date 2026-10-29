//ff:func feature=pack type=writer control=sequence
//ff:what TestOpenRuns — every run opens at its first entry; a missing run and a cut first entry fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenRuns(t *testing.T) {
	out := t.TempDir()
	dir := filepath.Join(out, streamDir)
	os.MkdirAll(dir, 0o755)
	os.WriteFile(filepath.Join(dir, "run-a"), appendHEntry(nil, hEntry{loc: Loc{Off: 5}}), 0o644)
	os.WriteFile(filepath.Join(dir, "run-cut"), make([]byte, 3), 0o644)
	rs, err := openRuns(out, []string{"run-a"})
	if err != nil || len(rs) != 1 || !rs[0].ok || rs[0].cur.loc.Off != 5 {
		t.Fatal(err)
	}
	closeRuns(rs)
	if _, err := openRuns(out, []string{"run-a", "run-missing"}); err == nil {
		t.Fatal("missing run: want error")
	}
	if _, err := openRuns(out, []string{"run-cut"}); err == nil {
		t.Fatal("cut run: want error")
	}
}
