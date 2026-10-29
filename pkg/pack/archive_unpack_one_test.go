//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveUnpackOne — one file with its folders; read and folder errors fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestArchiveUnpackOne(t *testing.T) {
	a, _ := testArchive(t)
	base := t.TempDir()
	if err := a.unpackOne(base, "0x0000/documents.geul"); err != nil {
		t.Fatal(err)
	}
	if err := a.unpackOne(base, "none"); err == nil {
		t.Fatal("read")
	}
	os.WriteFile(filepath.Join(base, "f"), nil, 0o644)
	if err := a.unpackOne(filepath.Join(base, "f"), "issuers.geul"); err == nil {
		t.Fatal("folder")
	}
}
