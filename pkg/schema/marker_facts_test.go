//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestMarkerFacts — UNKNOWN and NONE pass as Object/Value; markers in Predicate, Subject, Property or Target, ABSENT or a reserved marker as Object fail E7
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMarkerFacts(t *testing.T) {
	st := func(ws ...core.Word) statement {
		x, _, err := core.DecodeRecord(ws)
		s, ok := splitStatement(recOf(x))
		if err != nil || !ok {
			t.Fatalf("%04X", ws)
		}
		return s
	}
	for i, s := range []statement{
		st(0x140A, 0x0088, 1, 0x00A8, 40, 0x00A0, 30, core.NoValue, 0),
		st(0x138B, 0x0088, 2, 0x00A8, 582, 0x0088, 1, core.UnknownValue),
		st(0x148A, 0x0088, 3, 0x00A8, 31, 0x00A0, 30, 0x00A0, 5, 0),
	} {
		if fs := markerFacts(s); len(fs) != 0 {
			t.Fatalf("ok %d: %+v", i, fs)
		}
	}
	for want, s := range map[string]statement{
		"Predicate": st(0x140A, 0x0088, 1, core.UnknownValue, 0x00A0, 30, 0x00A0, 5, 0),
		"Subject":   st(0x140A, 0x0088, 1, 0x00A8, 31, core.NoValue, 0x00A0, 5, 0),
		"Object":    st(0x140A, 0x0088, 1, 0x00A8, 31, 0x00A0, 30, core.Absent, 0),
		"Value":     st(0x138B, 0x0088, 2, 0x00A8, 582, 0x0088, 1, 0x0005),
		"Property":  st(0x138B, 0x0088, 2, core.Absent, 0x0088, 1, 0x00A0, 5),
	} {
		fs := markerFacts(s)
		if len(fs) != 1 || fs[0].Rule != "E7" || fs[0].Where.Part != want {
			t.Fatalf("%s: %+v", want, fs)
		}
	}
}
