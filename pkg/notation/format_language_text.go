//ff:func feature=notation type=decoder control=sequence
//ff:what formatLanguageText — a UTF8-LANGUAGE-TEXT with a BCP 47 language as "text"@tag; false for another language code
package notation

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatLanguageText writes the text quoted and its LANGUAGE-BCP47-CODE tag
// after `@`. A short language code has no short form (false).
func formatLanguageText(r core.Record) (string, bool) {
	lang, text, err := core.DecodeLanguageText(r)
	if err != nil || lang.Kind != core.KindValue || lang.Codec&^1 != core.CodecLanguageBCP47Even {
		return "", false
	}
	tag, err := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even)
	return strconv.Quote(string(text)) + "@" + string(tag), err == nil
}
