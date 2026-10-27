//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirReadList — absent for an unpublished document, the list once published, errors for a bad list or an unreadable file
package resolve

import (
	"os"
	"testing"
)

func TestDirReadList(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := sampleKey(1, 0).DocKey()
	if lg, err := d.readList(k, k.Doc); err != nil || lg.Present {
		t.Fatalf("none: %+v %v", lg, err)
	}
	if err := d.write(listPath(k), listDocOf(1, VersionEntry{Version: 1})); err != nil {
		t.Fatal(err)
	}
	if lg, err := d.readList(k, k.Doc); err != nil || !lg.Present || len(lg.Entries) != 1 || len(lg.Records) != 1 {
		t.Fatalf("list: %+v %v", lg, err)
	}
	if err := d.write(listPath(k), []byte("junk")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.readList(k, k.Doc); err == nil {
		t.Fatal("bad list: want error")
	}
	if err := os.Remove(d.file(listPath(k))); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(d.file(listPath(k)), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := d.readList(k, k.Doc); err == nil {
		t.Fatal("unreadable: want error")
	}
}
