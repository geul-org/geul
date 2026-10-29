//ff:func feature=pack type=helper control=sequence
//ff:what TestFindDict — by ID, false when absent
package pack

import "testing"

func TestFindDict(t *testing.T) {
	ds := []dictEntry{{id: 1}, {id: 40000, loc: Loc{Off: 9}}}
	if d, ok := findDict(ds, 40000); !ok || d.loc.Off != 9 {
		t.Fatal(d)
	}
	if _, ok := findDict(ds, 2); ok {
		t.Fatal("absent")
	}
}
