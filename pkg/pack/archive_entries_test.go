//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveEntries — entries carry their full path and Archive; a file's Info has its real size
package pack

import "testing"

func TestArchiveEntries(t *testing.T) {
	a, _ := testArchive(t)
	idx, _ := a.index()
	es := a.entries(idx, ".")
	if len(es) != 3 || es[2].path != "issuers.geul" || es[2].a != a {
		t.Fatal(es)
	}
	fi, err := es[2].Info()
	if err != nil || fi.Size() != int64(len(fixtureFile(t, "issuers.geul"))) {
		t.Fatal(err)
	}
}
