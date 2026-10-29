//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveStat — file size and directory flag; absent is fs.ErrNotExist
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveStat(t *testing.T) {
	a, _ := testArchive(t)
	if fi, err := a.Stat("issuers.geul"); err != nil || fi.Size() != int64(len(fixtureFile(t, "issuers.geul"))) {
		t.Fatal(err)
	}
	if fi, err := fs.Stat(a, "0x0000"); err != nil || !fi.IsDir() {
		t.Fatal(err)
	}
	if _, err := a.Stat("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
