//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamSpill — nothing held writes nothing; held keys go sorted to a run file and the buffer empties; a missing working folder fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamSpill(t *testing.T) {
	s := streamTestNew(t, nil)
	if err := s.spill(); err != nil || len(s.st.Runs) != 0 {
		t.Fatal("empty spill")
	}
	var a, b hEntry
	a.key[0], b.key[0] = 2, 1
	s.hbuf = []hEntry{a, b}
	if err := s.spill(); err != nil || len(s.st.Runs) != 1 || s.st.RunCount != 1 || len(s.hbuf) != 0 {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(filepath.Join(s.out, streamDir, s.st.Runs[0]))
	if len(raw) != 2*hEntrySize || raw[0] != 1 || raw[hEntrySize] != 2 {
		t.Fatalf("run %x", raw[:1])
	}
	s.hbuf = []hEntry{a}
	os.RemoveAll(filepath.Join(s.out, streamDir))
	if err := s.spill(); err == nil {
		t.Fatal("no working folder: want error")
	}
}
