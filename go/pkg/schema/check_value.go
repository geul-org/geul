//ff:func feature=schema type=rule control=selection
//ff:what checkValue — V1~V5 for one Value record
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkValue applies the Value rules to a direct-code Value. ISSUER formats are
// left to layer 2 (F5 already checked their extension TAG).
func checkValue(r rec) []Fact {
	c := r.R.Codec
	switch {
	case c >= core.CodecIssuer1Word:
		return nil
	case c == 0x00:
		return checkBool(r)
	case c == 0x01 || c == core.CodecUnsignedInteger:
		return checkInteger(r, r.R.Payload, c == 0x01, "Payload")
	case c == 0x03:
		return checkWidth(r, 1, 2, 4, 8)
	case c == core.CodecSignedDecimal || c == core.CodecUnsignedDecimal:
		return checkDecimal(r, c == core.CodecSignedDecimal)
	case c == 0x21:
		return checkCivilTime(r)
	case fixedValueWidth[c] > 0:
		return checkWidth(r, fixedValueWidth[c])
	case c >= core.CodecUTF8TextEven && c <= 0x47, c == core.CodecEDTFEven, c == core.CodecEDTFEven|1:
		return checkBytes(r)
	default:
		return []Fact{factAt("V1", Warn, r, "Codec", "배정된 직접 코드", fmt.Sprintf("0x%02X (예약)", c))}
	}
}
