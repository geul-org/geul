//ff:func feature=resolve type=resolver control=sequence
//ff:what TestListNext — the NEXT-LOCAL-ID among HEAD records; none or a bad one gives 0
package resolve

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestListNext(t *testing.T) {
	if n := listNext([]core.Record{core.EncodeSpecVersion(2), core.EncodeNextLocalID(9)}); n != 9 {
		t.Fatalf("got %d", n)
	}
	if listNext([]core.Record{core.EncodeSpecVersion(2), core.EncodeNextLocalID(0)}) != 0 || listNext(nil) != 0 {
		t.Fatal("none: want 0")
	}
}
