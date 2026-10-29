//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveIsDir — the top, folders holding paths; files and absent names are not
package pack

import "testing"

func TestArchiveIsDir(t *testing.T) {
	a, _ := testArchive(t)
	if !a.isDir(".") || !a.isDir("0x0000") || !a.isDir("h") || a.isDir("issuers.geul") || a.isDir("zz") {
		t.Fatal("isDir")
	}
}
