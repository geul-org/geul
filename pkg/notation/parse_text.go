//ff:func feature=notation type=encoder control=selection
//ff:what parseText — "text" as UTF8-TEXT, "text"@tag as UTF8-LANGUAGE-TEXT with a LANGUAGE-BCP47-CODE, "text"@@tag or @@0xNNNN with a LANGUAGE-SHORT-CODE
package notation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseText reads a quoted string token: alone a UTF8-TEXT, followed by
// `@tag` a UTF8-LANGUAGE-TEXT whose language is that LANGUAGE-BCP47-CODE, and
// followed by `@@` a UTF8-LANGUAGE-TEXT whose language is a
// LANGUAGE-SHORT-CODE (shortLanguage: a registry tag or the code in hex).
func parseText(tok string, ctx Context) (core.Record, error) {
	end := strings.LastIndex(tok, `"`)
	s, err := strconv.Unquote(tok[:end+1])
	if err != nil {
		return core.Record{}, fmt.Errorf("geul/notation: %s is not a quoted string", tok)
	}
	rest := tok[end+1:]
	switch {
	case rest == "":
		return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s)), nil
	case strings.HasPrefix(rest, "@@"):
		lang, err := shortLanguage(rest[2:], ctx)
		return core.EncodeLanguageText(lang, []byte(s)), err
	case strings.HasPrefix(rest, "@") && len(rest) > 1:
		lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(rest[1:]))
		return core.EncodeLanguageText(lang, []byte(s)), nil
	default:
		return core.Record{}, fmt.Errorf("geul/notation: %q after a quoted string", rest)
	}
}
