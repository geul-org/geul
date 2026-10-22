//ff:func feature=geulmd type=encoder control=selection
//ff:what renderValue — a typed literal for the Value codecs GEULMD knows; false for any other
package geulmd

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// renderValue renders a Value record as a typed literal: UTF8-TEXT as a quoted
// string, a BCP 47 tag as `lang"…"`, UNSIGNED-INTEGER as `uint(n)` and the
// decimals as `dec(CeE)` / `udec(CeE)` (coefficient and exponent kept, so the
// quantum survives). ok is false for any other codec or a malformed payload.
func renderValue(r core.Record) (string, bool) {
	switch r.Codec &^ 1 {
	case core.CodecUTF8TextEven:
		b, err := core.DecodeBytesValue(r, core.CodecUTF8TextEven)
		return strconv.Quote(string(b)), err == nil
	case core.CodecLanguageBCP47Even:
		b, err := core.DecodeBytesValue(r, core.CodecLanguageBCP47Even)
		return "lang" + strconv.Quote(string(b)), err == nil
	case core.CodecUnsignedInteger:
		n, err := core.DecodeUintValue(r)
		return "uint(" + strconv.FormatUint(n, 10) + ")", err == nil
	case core.CodecSignedDecimal, core.CodecUnsignedDecimal:
		e, c, err := core.DecodeDecimalValue(r)
		name := "dec("
		if r.Codec == core.CodecUnsignedDecimal {
			name = "udec("
		}
		return name + strconv.FormatInt(c, 10) + "e" + strconv.FormatInt(int64(e), 10) + ")", err == nil
	default:
		return "", false
	}
}
