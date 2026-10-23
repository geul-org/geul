//ff:func feature=resolve type=rule control=sequence
//ff:what TestMarkVersion — DEPRECATION and TOMBSTONE flag their version; an unpublished version fails
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestMarkVersion(t *testing.T) {
	lg := &versionLog{Entries: []VersionEntry{{Version: 1}, {Version: 2}}}
	if markVersion(lg, core.Publication{Codec: core.EdgeDeprecation, Version: 1}) != nil || markVersion(lg, core.Publication{Codec: core.EdgeTombstone, Version: 2}) != nil {
		t.Fatal("mark")
	}
	if !lg.Entries[0].Deprecated || lg.Entries[0].Deleted || !lg.Entries[1].Deleted || lg.Entries[1].Deprecated {
		t.Fatalf("flags: %+v", lg.Entries)
	}
	if markVersion(lg, core.Publication{Codec: core.EdgeDeprecation, Version: 9}) == nil {
		t.Fatal("unpublished: want error")
	}
}
