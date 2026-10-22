//ff:func feature=inspect type=decoder control=selection
//ff:what parseValue — a Value record of a known codec as Value; false for any other codec or a malformed payload
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseValue decodes the Value codecs this library knows: UTF8-TEXT and BCP 47
// tags (either pair member), UNSIGNED-INTEGER and the two decimals. ok is
// false for any other codec or a malformed payload.
func parseValue(r core.Record) (Value, bool) {
	v := Value{Codec: r.Codec}
	var err error
	switch even := r.Codec &^ 1; even {
	case core.CodecUTF8TextEven, core.CodecLanguageBCP47Even:
		var b []byte
		b, err = core.DecodeBytesValue(r, even)
		v.Text = string(b)
	case core.CodecUnsignedInteger:
		v.Uint, err = core.DecodeUintValue(r)
	case core.CodecSignedDecimal:
		v.Exp, v.Coef, err = core.DecodeDecimalValue(r)
	default:
		return Value{}, false
	}
	return v, err == nil
}
