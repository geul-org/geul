//ff:func feature=resolve type=helper control=sequence
//ff:what TestDeletedIn — a deleted version is found; a live, a missing and version 0 are not
package resolve

import "testing"

func TestDeletedIn(t *testing.T) {
	list := []VersionEntry{{Version: 1, Deleted: true}, {Version: 2}}
	if !deletedIn(list, 1) || deletedIn(list, 2) || deletedIn(list, 3) || deletedIn(list, 0) {
		t.Fatal("deletedIn")
	}
}
