//ff:func feature=notation type=decoder control=selection
//ff:what formatValue — a Value record as a token in its short form; false when the notation has none
package notation

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatValue writes the Value codecs with a short form (package doc): byte
// strings, language text with a BCP 47 tag, the decimals, unsigned integers,
// the two time codecs, QUANTITY and the digests. Others are false.
func formatValue(r core.Record) (string, bool) {
	name := core.CodecName(core.KindValue, r.Codec)
	switch even := r.Codec &^ 1; {
	case even == core.CodecUTF8TextEven || even == codecURIEven || even == core.CodecEDTFEven || even == core.CodecGEDCOMDateEven || even == core.CodecLanguageBCP47Even:
		b, err := core.DecodeBytesValue(r, even)
		return formatTextValue(even, string(b)), err == nil
	case even == core.CodecUTF8LanguageTextEven:
		return formatLanguageText(r)
	case r.Codec == core.CodecSignedDecimal || r.Codec == core.CodecUnsignedDecimal:
		return formatDecimal(r)
	case r.Codec == core.CodecUnsignedInteger:
		n, err := core.DecodeUintValue(r)
		return name + "(" + strconv.FormatUint(n, 10) + ")", err == nil
	case r.Codec == core.CodecGregorianDate || r.Codec == core.CodecPosixSeconds:
		return formatTime(r)
	case r.Codec == core.CodecQuantity:
		return formatQuantity(r)
	case r.Codec == core.CodecSHA256Digest || r.Codec == core.CodecSHA1Digest:
		return name + "(" + hexToken(r.Payload) + ")", true
	default:
		return "", false
	}
}
