//ff:func feature=notation type=decoder control=selection
//ff:what formatLanguageText — a UTF8-LANGUAGE-TEXT as "text"@tag (LANGUAGE-BCP47-CODE) or "text"@@0xNNNN (LANGUAGE-SHORT-CODE, named later by FormatEntry)
package notation

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatLanguageText writes the text quoted and its language after it: a
// LANGUAGE-BCP47-CODE tag after `@`, a LANGUAGE-SHORT-CODE after `@@` as its
// code in hex (FormatEntry puts the registry tag in its place,
// namedLanguageToken). A record that does not decode is false.
func formatLanguageText(r core.Record) (string, bool) {
	lang, text, err := core.DecodeLanguageText(r)
	switch {
	case err != nil || lang.Kind != core.KindValue:
		return "", false
	case lang.Codec == codecLanguageShort && len(lang.Payload) == 1:
		return strconv.Quote(string(text)) + "@@" + core.FormatID(lang.Payload), true
	case lang.Codec&^1 == core.CodecLanguageBCP47Even:
		tag, err := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even)
		return strconv.Quote(string(text)) + "@" + string(tag), err == nil
	default:
		return "", false
	}
}
