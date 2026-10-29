//ff:func feature=pack type=writer control=sequence
//ff:what TestNewStream — a new set gets generation 1 and created_at from the clock when not given; samples train a dictionary stored first; an existing set, an unfinished stream and an unmakeable folder are refused
package pack

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewStream(t *testing.T) {
	o := streamTestOptions()
	o.Created, o.Build.Now = 0, func() time.Time { return time.Unix(1_900_000_000, 0) }
	out := filepath.Join(t.TempDir(), "set")
	s, err := NewStream(out, o)
	if err != nil || s.st.Gen != 1 || s.st.Created != 1_900_000_000 || s.st.Dict.ID != FirstDictID || s.dict == nil || s.st.Data.Count != 1 {
		t.Fatalf("%+v %v", s.st.Dict, err)
	}
	s.Checkpoint(nil)
	if _, err := NewStream(out, o); err == nil {
		t.Fatal("unfinished stream: want error")
	}
	done := filepath.Join(t.TempDir(), "done")
	os.MkdirAll(done, 0o755)
	os.WriteFile(filepath.Join(done, latestFile), []byte("1 x\n"), 0o644)
	if _, err := NewStream(done, o); err == nil {
		t.Fatal("existing set: want error")
	}
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	if _, err := NewStream(filepath.Join(file, "set"), o); err == nil {
		t.Fatal("under a file: want error")
	}
	o.Samples = nil
	if s, err := NewStream(filepath.Join(t.TempDir(), "plain"), o); err != nil || s.dict != nil || s.st.Data.Num != -1 {
		t.Fatalf("no samples: %v", err)
	}
}
