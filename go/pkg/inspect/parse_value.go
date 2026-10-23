//ff:func feature=inspect type=decoder control=selection
//ff:what parseValue — a Value record of a known codec as Value; false for any other codec or a malformed payload
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseValue decodes the Value codecs this library knows: UTF8-TEXT and BCP 47
// tags (either pair member), UNSIGNED-INTEGER, the two decimals, the two time
// codecs and SHA-256-DIGEST. ok is
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
	case core.CodecGregorianDate, core.CodecPosixSeconds:
		v.Time, err = core.DecodeTimeValue(r)
	case core.CodecSHA256Digest &^ 1:
		v.Text = core.FormatID(r.Payload)
		return v, r.Codec == core.CodecSHA256Digest && len(r.Payload) == 16
	default:
		return Value{}, false
	}
	return v, err == nil
}
