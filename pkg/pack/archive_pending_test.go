//ff:func feature=pack type=writer control=sequence
//ff:what TestArchivePending — true while changes are staged
package pack

import "testing"

func TestArchivePending(t *testing.T) {
	a, _ := testArchive(t)
	if a.Pending() {
		t.Fatal("clean")
	}
	a.WriteFile("x", nil)
	if !a.Pending() {
		t.Fatal("staged")
	}
}
