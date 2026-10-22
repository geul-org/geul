//ff:func feature=geulmd type=encoder control=selection
//ff:what renderValue — a typed literal for an interpreted Value
package geulmd

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// renderValue renders a Value as a typed literal: UTF8-TEXT as a quoted
// string, a BCP 47 tag as `lang"…"`, UNSIGNED-INTEGER as `uint(n)` and the
// decimals as `dec(CeE)` / `udec(CeE)` (coefficient and exponent kept, so the
// quantum survives).
func renderValue(v inspect.Value) string {
	switch v.Codec &^ 1 {
	case core.CodecUTF8TextEven:
		return strconv.Quote(v.Text)
	case core.CodecLanguageBCP47Even:
		return "lang" + strconv.Quote(v.Text)
	case core.CodecUnsignedInteger:
		return "uint(" + strconv.FormatUint(v.Uint, 10) + ")"
	default:
		name := "dec("
		if v.Codec == core.CodecUnsignedDecimal {
			name = "udec("
		}
		return name + strconv.FormatInt(v.Coef, 10) + "e" + strconv.FormatInt(int64(v.Exp), 10) + ")"
	}
}
