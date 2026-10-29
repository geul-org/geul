//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchivePaths — a copy of the sorted listing, staged changes included
package pack

import "testing"

func TestArchivePaths(t *testing.T) {
	a, _ := testArchive(t)
	a.WriteFile("zz.geul", nil)
	ps, err := a.Paths()
	if err != nil || len(ps) != 19 || ps[18] != "zz.geul" {
		t.Fatal(ps, err)
	}
	ps[0] = "x"
	if again, _ := a.Paths(); again[0] == "x" {
		t.Fatal("not a copy")
	}
}
