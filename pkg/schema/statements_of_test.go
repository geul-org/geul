//ff:func feature=schema type=decoder control=sequence
//ff:what TestStatementsOf — only TRIPLEs and QUALIFIERs are kept, in order
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStatementsOf(t *testing.T) {
	tr, _, _ := core.DecodeRecord([]core.Word{0x140A, 0x0088, 7, 0x00A8, 40, 0x00A0, 42, core.NoValue, 0})
	q, _, _ := core.DecodeRecord([]core.Word{0x138B, 0x0088, 8, 0x00A8, 582, 0x0088, 7, core.UnknownValue})
	sts := statementsOf([]rec{recOf(tr), recOf(core.EncodeLocalRef(3)), recOf(q)})
	if len(sts) != 2 || sts[0].ID != 7 || sts[1].ID != 8 {
		t.Fatalf("%+v", sts)
	}
}
