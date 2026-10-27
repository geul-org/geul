//ff:func feature=resolve type=helper control=sequence
//ff:what TestFindHash — found by content hash, absent otherwise
package resolve

import (
	"testing"
)

func TestFindHash(t *testing.T) {
	list := []VersionEntry{{Version: 1, Hash: [32]byte{1}}, {Version: 2, Hash: [32]byte{2}}}
	if e, ok := findHash(list, [32]byte{2}); !ok || e.Version != 2 {
		t.Fatalf("found: %+v %v", e, ok)
	}
	if _, ok := findHash(list, [32]byte{3}); ok {
		t.Fatal("absent: want false")
	}
}
