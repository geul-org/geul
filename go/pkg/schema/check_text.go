//ff:func feature=schema type=rule control=selection
//ff:what checkText — V5: UTF-8 validity and no BOM for text, non-empty ASCII for URI, language tags, EDTF and GEDCOM-DATE; V6 for EDTF
package schema

import (
	"bytes"
	"unicode/utf8"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkText applies V5 to the bytes of a byte-string Value whose even codec is
// even (0x40 UTF8-TEXT, 0x44 URI, 0x46 LANGUAGE-BCP47-CODE,
// 0x24 EDTF, 0x26 GEDCOM-DATE), then V6 to EDTF. The EDTF and GEDCOM date
// grammars are not checked yet (README 미결정).
func checkText(r rec, even uint8, b []byte) []Fact {
	text := even == core.CodecUTF8TextEven
	switch {
	case text && !utf8.Valid(b):
		return []Fact{factAt("V5", Error, r, "Payload", "well-formed UTF-8", "invalid UTF-8")}
	case text && bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}):
		return []Fact{factAt("V5", Error, r, "Payload", "선두 BOM 없음", "EF BB BF")}
	case text:
		return nil
	case len(b) == 0:
		return []Fact{factAt("V5", Error, r, "Payload", "빈 문자열이 아님", "0 byte")}
	case even == core.CodecGEDCOMDateEven && !gedcomASCII(b):
		return []Fact{factAt("V5", Error, r, "Payload", "ASCII (0x20~0x7E), 앞뒤 공백 없음", "다른 byte 또는 앞뒤 공백")}
	case even == core.CodecGEDCOMDateEven:
		return nil
	case !printableASCII(b):
		return []Fact{factAt("V5", Error, r, "Payload", "ASCII (0x21~0x7E)", "다른 byte 포함")}
	case even == core.CodecEDTFEven && edtfPlainDay(b):
		return []Fact{factAt("V6", Error, r, "Payload", "GREGORIAN-DATE로 쓰는 날짜가 아님", string(b))}
	default:
		return nil
	}
}
