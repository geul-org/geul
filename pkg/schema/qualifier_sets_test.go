//ff:func feature=schema type=helper control=sequence
//ff:what TestQualifierSets — QUALIFIERs grouped by Target, order-free; TRIPLEs and a non-LOCAL Target are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQualifierSets(t *testing.T) {
	st := func(ws ...core.Word) statement {
		x, _, _ := core.DecodeRecord(ws)
		s, ok := splitStatement(recOf(x))
		if !ok {
			t.Fatalf("%04X", ws)
		}
		return s
	}
	a := st(0x138B, 0x0088, 8, 0x00A8, 582, 0x0088, 7, core.UnknownValue)
	b := st(0x138B, 0x0088, 9, 0x00A8, 580, 0x0088, 7, core.NoValue)
	tr := st(0x140A, 0x0088, 7, 0x00A8, 40, 0x00A0, 42, core.NoValue, 0)
	bad := st(0x138B, 0x0088, 10, 0x00A8, 582, 0x00A0, 7, core.UnknownValue)
	m1, m2 := qualifierSets([]statement{a, b, tr, bad}), qualifierSets([]statement{b, a})
	if len(m1) != 1 || m1[7] == "" || m1[7] != m2[7] {
		t.Fatalf("%v %v", m1, m2)
	}
}
