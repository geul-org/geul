//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveFetch — bytes from the serving set, false when none serves the path, broken sets fail
package pack

import (
	"os"
	"testing"
)

func TestArchiveFetch(t *testing.T) {
	a, _ := testArchive(t)
	if b, ok, err := a.fetch("0x0000/documents.geul"); !ok || err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	if _, ok, err := a.fetch("0x0000/none.geul"); ok || err != nil {
		t.Fatal(err)
	}
	os.Remove(objPath(t, a.sets[0].r, a.sets[0].dir, kindData))
	if _, _, err := a.fetch("0x0000/documents.geul"); err == nil {
		t.Fatal("broken")
	}
}
