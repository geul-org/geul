//ff:func feature=resolve type=helper control=sequence
//ff:what TestPickVersion — exact version, latest skips deprecated and deleted, deleted, missing and all-deprecated fail
package resolve

import (
	"testing"
)

func TestPickVersion(t *testing.T) {
	list := []VersionEntry{{Version: 1}, {Version: 2}, {Version: 3, Deprecated: true}, {Version: 4, Deleted: true}}
	if e, ok := pickVersion(list, 3); !ok || e.Version != 3 {
		t.Fatalf("exact: %+v %v", e, ok)
	}
	if e, ok := pickVersion(list, 0); !ok || e.Version != 2 {
		t.Fatalf("latest: %+v %v", e, ok)
	}
	if _, ok := pickVersion(list, 4); ok {
		t.Fatal("deleted: want false")
	}
	if _, ok := pickVersion(list, 9); ok {
		t.Fatal("missing: want false")
	}
	if _, ok := pickVersion(list[2:], 0); ok {
		t.Fatal("all deprecated: want false")
	}
}
