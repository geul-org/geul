//ff:func feature=pack type=writer control=sequence
//ff:what TestHashOwners — content hashes of Issuer documents map to their Issuer; other paths and non-documents are skipped; unreadable files fail
package pack

import (
	"path/filepath"
	"testing"
)

func TestHashOwners(t *testing.T) {
	base := filepath.Join(fixture(t), ".well-known", "geul")
	paths, _ := listTree(filepath.Dir(filepath.Dir(base)))
	hs, err := hashOwners(base, append(paths, "0x0000/none.geul")[:len(paths)])
	if err != nil || len(hs) != 9 {
		t.Fatal(len(hs), err)
	}
	h, _ := contentHash(fixtureFile(t, "0x0000/0x0000000000000845/versions.geul"))
	if hs[string(h[:])] != "0x0000" {
		t.Fatal("owner")
	}
	if _, err := hashOwners(base, []string{"0x0000/none.geul"}); err == nil {
		t.Fatal("unreadable")
	}
}
