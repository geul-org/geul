//ff:func feature=notation type=decoder control=sequence
//ff:what formatInteger — a SIGNED-INTEGER or UNSIGNED-INTEGER of any size as NAME(digits)
package notation

import "github.com/geul-org/gogeul/pkg/core"

// formatInteger writes an integer Value (core.DecodeIntegerBig) as its codec
// name and decimal digits: SIGNED-INTEGER(-3), UNSIGNED-INTEGER(5). A
// non-minimal payload is false (raw words instead).
func formatInteger(r core.Record) (string, bool) {
	n, err := core.DecodeIntegerBig(r)
	if err != nil {
		return "", false
	}
	return core.CodecName(core.KindValue, r.Codec) + "(" + n.String() + ")", true
}
