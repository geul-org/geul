//ff:func feature=schema type=decoder control=sequence
//ff:what TestEdgeIDs — the Edge Ids of a BODY; records without one are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeIDs(t *testing.T) {
	tr, _, _ := core.DecodeRecord([]core.Word{0x148A, 0x0088, 5, 0x00A8, 31, 0x00A0, 42, 0x00A0, 5, 0})
	ids := edgeIDs([]rec{recOf(tr), recOf(core.EncodeLocalRef(7))})
	if len(ids) != 1 || !ids[5] {
		t.Fatalf("%v", ids)
	}
}
