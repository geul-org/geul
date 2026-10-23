//ff:func feature=schema type=rule control=selection
//ff:what checkText — V5: UTF-8 validity and no BOM for text, non-empty ASCII for URI and language tags
package schema

import (
	"bytes"
	"unicode/utf8"
)

// checkText applies V5 to the bytes of a byte-string Value whose even codec is
// even (0x40 UTF8-TEXT, 0x42 UTF8-NFC-TEXT, 0x44 URI, 0x46 LANGUAGE-BCP47-CODE).
// NFC itself is not checked yet (README 미결정).
func checkText(r rec, even uint8, b []byte) []Fact {
	switch {
	case even <= 0x42 && !utf8.Valid(b):
		return []Fact{factAt("V5", Error, r, "Payload", "well-formed UTF-8", "invalid UTF-8")}
	case even <= 0x42 && bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}):
		return []Fact{factAt("V5", Error, r, "Payload", "선두 BOM 없음", "EF BB BF")}
	case even <= 0x42:
		return nil
	case len(b) == 0:
		return []Fact{factAt("V5", Error, r, "Payload", "빈 문자열이 아님", "0 byte")}
	case !printableASCII(b):
		return []Fact{factAt("V5", Error, r, "Payload", "ASCII (0x21~0x7E)", "다른 byte 포함")}
	default:
		return nil
	}
}
