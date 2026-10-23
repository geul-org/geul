//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirReadList — none for an unpublished document, the list once published, errors for a bad list or unreadable file
package resolve

import (
	"os"
	"testing"
)

func TestDirReadList(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k := sampleKey(1, 0).DocKey()
	if l, err := d.readList(k); err != nil || l != nil {
		t.Fatalf("none: %+v %v", l, err)
	}
	if err := d.write(listPath(k), []byte(formatVersionList([]VersionEntry{{Version: 1}}))); err != nil {
		t.Fatal(err)
	}
	if l, err := d.readList(k); err != nil || len(l) != 1 {
		t.Fatalf("list: %+v %v", l, err)
	}
	if err := d.write(listPath(k), []byte("junk")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.readList(k); err == nil {
		t.Fatal("bad list: want error")
	}
	if err := os.MkdirAll(d.file(listPath(k)+"x"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := d.readList(DocKey{Issuer: k.Issuer, Doc: k.Doc + "/versions"}); err == nil {
		t.Fatal("unreadable: want error")
	}
}
