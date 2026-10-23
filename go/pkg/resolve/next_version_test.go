//ff:func feature=resolve type=helper control=sequence
//ff:what TestNextVersion — 1 for an empty list, last+1, exhausted at the 2-word maximum
package resolve

import (
	"math"
	"testing"
)

func TestNextVersion(t *testing.T) {
	if v, ok := nextVersion(nil); !ok || v != 1 {
		t.Fatalf("empty: %d %v", v, ok)
	}
	if v, ok := nextVersion([]VersionEntry{{Version: 4}}); !ok || v != 5 {
		t.Fatalf("next: %d %v", v, ok)
	}
	if _, ok := nextVersion([]VersionEntry{{Version: math.MaxUint32}}); ok {
		t.Fatal("exhausted: want false")
	}
}
