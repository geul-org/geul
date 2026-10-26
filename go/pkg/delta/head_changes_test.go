//ff:func feature=delta type=decoder control=sequence
//ff:what TestHeadChanges — equal records match once; duplicates and leftovers are removed and added
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHeadChanges(t *testing.T) {
	a, b := core.EncodeSpecVersion(2), core.EncodeNextLocalID(3)
	cs := headChanges([]core.Record{a, a, b}, []core.Record{a, core.EncodeNextLocalID(4)})
	if len(cs) != 3 || cs[0].Op != '-' || cs[1].Op != '-' || cs[2].Op != '+' || !cs[0].Head {
		t.Fatalf("%+v", cs)
	}
	if cs := headChanges([]core.Record{a}, []core.Record{a}); len(cs) != 0 {
		t.Fatal("same head: want none")
	}
}
