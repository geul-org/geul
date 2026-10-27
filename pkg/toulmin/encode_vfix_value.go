//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeVFixValue — a verdict/qualifier VFix as a trimmed SIGNED-DECIMAL Value record
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// EncodeVFixValue writes v (scale 10000) as a SIGNED-DECIMAL Value record with
// the trailing zeros trimmed: 0.8 is C=8, E=-1 (GEUL t002 TOULMIN 수치).
func EncodeVFixValue(v VFix) core.Record {
	e, c := core.TrimDecimal(-4, int64(v))
	r, _ := core.EncodeDecimalValue(true, e, c)
	return r
}
