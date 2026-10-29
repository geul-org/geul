//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamSaveState — the state is replaced in one rename; a missing working folder and a state path that is a folder fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamSaveState(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Meta = []byte("x")
	if err := s.saveState(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(s.out, streamDir, streamStateFile+".tmp")); err == nil {
		t.Fatal("temp file left")
	}
	os.Remove(filepath.Join(s.out, streamDir, streamStateFile))
	os.MkdirAll(filepath.Join(s.out, streamDir, streamStateFile, "x"), 0o755)
	if err := s.saveState(); err == nil {
		t.Fatal("state path is a folder: want error")
	}
	os.RemoveAll(filepath.Join(s.out, streamDir))
	if err := s.saveState(); err == nil {
		t.Fatal("no working folder: want error")
	}
}
