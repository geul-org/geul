//ff:func feature=resolve type=resolver control=sequence
//ff:what TestLatestStored — the highest version not deleted, a deprecated one included; none when all are deleted
package resolve

import "testing"

func TestLatestStored(t *testing.T) {
	es := []VersionEntry{{Version: 1}, {Version: 2, Deprecated: true}, {Version: 3, Deleted: true}}
	if e, ok := latestStored(es); !ok || e.Version != 2 {
		t.Fatalf("%+v %v", e, ok)
	}
	if _, ok := latestStored([]VersionEntry{{Version: 1, Deleted: true}}); ok {
		t.Fatal("all deleted: want none")
	}
}
