//ff:func feature=page type=helper control=selection
//ff:what valueText — a Value as readable text: quoted string, language tag, integer, time, digest or decimal
package page

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// valueText shows a Value: UTF8-TEXT quoted, UTF8-LANGUAGE-TEXT quoted with its
// language ("Paris"@fr), a URI, BCP 47 tag, EDTF or GEDCOM-DATE string bare, an integer in
// decimal, a BINARY-FLOAT in shortest form, a coordinate by coordinateText, a quantity by quantityText, a date or POSIX time in ISO 8601 (UTC), a digest in hex, and a decimal as a number followed by its exact C e E form.
func valueText(v inspect.Value) string {
	if v.Codec == core.CodecBinaryFloat {
		return strconv.FormatFloat(v.Float, 'g', -1, 64)
	}
	if v.Codec == core.CodecQuantity {
		return quantityText(v)
	}
	switch v.Codec &^ 1 {
	case core.CodecUTF8TextEven:
		return strconv.Quote(v.Text)
	case core.CodecUTF8LanguageTextEven:
		return strconv.Quote(v.Text) + "@" + v.Lang
	case 0x44, core.CodecLanguageBCP47Even, core.CodecEDTFEven, core.CodecGEDCOMDateEven:
		return v.Text
	case core.CodecCoordinateEarth:
		return coordinateText(v)
	case core.CodecUnsignedInteger:
		return strconv.FormatUint(v.Uint, 10)
	case core.CodecGregorianDate:
		return v.Time.Format(time.DateOnly)
	case core.CodecPosixSeconds:
		return v.Time.Format(time.RFC3339)
	case core.CodecSHA256Digest &^ 1, core.CodecSHA1Digest:
		return v.Text
	default:
		f := float64(v.Coef) * math.Pow10(int(v.Exp))
		return fmt.Sprintf("%s (%de%d)", strconv.FormatFloat(f, 'g', -1, 64), v.Coef, v.Exp)
	}
}
