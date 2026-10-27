//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what sampleLawTree — test fixture: document 5 amended three times with effective dates, document 6 importing it latest
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

// sampleLawTree publishes document 5 in three versions: v1 in effect from day
// 100, v2 from day 200 until day 300 (a sunset law), v3 from day 400. Document 6
// imports document 5 at latest.
func sampleLawTree(t *testing.T) Dir {
	t.Helper()
	d := Dir{Root: t.TempDir()}
	from, until := uint32(core.MetaEffectiveFrom), uint32(core.MetaEffectiveUntil)
	for _, b := range [][]byte{
		sampleEffectiveDoc(5, "v1", [2]uint32{from, 100}),
		sampleEffectiveDoc(5, "v2", [2]uint32{from, 200}, [2]uint32{until, 300}),
		sampleEffectiveDoc(5, "v3", [2]uint32{from, 400}),
		sampleDoc(6, "case", [3]uint32{1, 5, 0}),
	} {
		if _, err := d.Publish(b); err != nil {
			t.Fatal(err)
		}
	}
	return d
}
