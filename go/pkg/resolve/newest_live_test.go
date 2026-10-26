//ff:func feature=resolve type=resolver control=sequence
//ff:what TestNewestLive — the highest live version; deleted and deprecated ones are skipped; none gives false
package resolve

import "testing"

func TestNewestLive(t *testing.T) {
	es := []VersionEntry{{Version: 1}, {Version: 3, Deleted: true}, {Version: 2}, {Version: 4, Deprecated: true}}
	if e, ok := newestLive(es); !ok || e.Version != 2 {
		t.Fatalf("%+v %v", e, ok)
	}
	if _, ok := newestLive([]VersionEntry{{Version: 1, Deleted: true}}); ok {
		t.Fatal("none: want false")
	}
}
