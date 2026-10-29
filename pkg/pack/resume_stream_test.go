//ff:func feature=pack type=writer control=sequence
//ff:what TestResumeStream — the checkpoint comes back with its dictionary, open parts and meta; no state, a corrupt state and a lost part file fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResumeStream(t *testing.T) {
	out := filepath.Join(t.TempDir(), "set")
	s, _ := NewStream(out, streamTestOptions())
	streamTestPut(t, s, 1, 3)
	s.Checkpoint([]byte("three"))
	r, err := ResumeStream(out, streamTestOptions())
	if err != nil || r.Count() != 6 || string(r.Meta()) != "three" || r.dict == nil || r.data == nil || r.index == nil || r.inum["0x00010000"] != 0 {
		t.Fatalf("%v", err)
	}
	if _, err := ResumeStream(t.TempDir(), streamTestOptions()); err == nil {
		t.Fatal("no state: want error")
	}
	state := filepath.Join(out, streamDir, streamStateFile)
	good, _ := os.ReadFile(state)
	os.WriteFile(state, []byte("not gob"), 0o644)
	if _, err := ResumeStream(out, streamTestOptions()); err == nil {
		t.Fatal("corrupt state: want error")
	}
	os.WriteFile(state, good, 0o644)
	os.Remove(filepath.Join(out, streamDir, streamIndexPart))
	if _, err := ResumeStream(out, streamTestOptions()); err == nil {
		t.Fatal("lost index part: want error")
	}
	os.Remove(filepath.Join(out, streamDir, streamDataPart))
	if _, err := ResumeStream(out, streamTestOptions()); err == nil {
		t.Fatal("lost data part: want error")
	}
	plain := filepath.Join(t.TempDir(), "plain")
	o := streamTestOptions()
	o.Samples = nil
	p, _ := NewStream(plain, o)
	p.Checkpoint(nil)
	if r, err := ResumeStream(plain, o); err != nil || r.dict != nil || r.data != nil {
		t.Fatalf("empty stream: %v", err)
	}
}
