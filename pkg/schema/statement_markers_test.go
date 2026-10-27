//ff:func feature=schema type=rule control=sequence
//ff:what TestStatementMarkers — E7 over a BODY: a misplaced marker in a TRIPLE is reported, other records are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStatementMarkers(t *testing.T) {
	bad, _, _ := core.DecodeRecord([]core.Word{0x140A, 0x0088, 1, 0x00A8, 31, 0x00A0, 30, core.Absent, 0})
	good, _, _ := core.DecodeRecord([]core.Word{0x140A, 0x0088, 2, 0x00A8, 31, 0x00A0, 30, core.UnknownValue, 0})
	if got := rulesOf(statementMarkers([]rec{recOf(bad), recOf(good), recOf(core.EncodeLocalRef(1))})); got != "E7" {
		t.Fatalf("%q", got)
	}
}
