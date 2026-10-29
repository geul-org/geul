//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestMakeStreamDirs — o/, m/ and .stream are made; a folder under a file fails
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeStreamDirs(t *testing.T) {
	out := filepath.Join(t.TempDir(), "set")
	if err := makeStreamDirs(out); err != nil {
		t.Fatal(err)
	}
	for _, d := range []string{"o", "m", streamDir} {
		if fi, err := os.Stat(filepath.Join(out, d)); err != nil || !fi.IsDir() {
			t.Fatalf("%s: %v", d, err)
		}
	}
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	if err := makeStreamDirs(filepath.Join(file, "set")); err == nil {
		t.Fatal("under a file: want error")
	}
}
