//ff:func feature=pack type=writer control=sequence
//ff:what TestOpenPart — none open is nil; a part is cut back to the checkpoint; a missing part and one shorter than the checkpoint fail
package pack

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenPart(t *testing.T) {
	out := t.TempDir()
	os.MkdirAll(filepath.Join(out, streamDir), 0o755)
	if f, err := openPart(out, streamDataPart, streamOpen{Num: -1}); f != nil || err != nil {
		t.Fatal("none open")
	}
	if _, err := openPart(out, streamDataPart, streamOpen{Num: 0, Off: 4}); err == nil {
		t.Fatal("missing part: want error")
	}
	os.WriteFile(filepath.Join(out, streamDir, streamDataPart), []byte("0123456789"), 0o644)
	f, err := openPart(out, streamDataPart, streamOpen{Num: 0, Off: 4})
	if err != nil {
		t.Fatal(err)
	}
	if fi, _ := f.Stat(); fi.Size() != 4 {
		t.Fatalf("size %d", fi.Size())
	}
	if _, err := openPart(out, streamDataPart, streamOpen{Num: 0, Off: 8}); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("short part: %v", err)
	}
}
