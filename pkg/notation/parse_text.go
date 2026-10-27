//ff:func feature=notation type=encoder control=selection
//ff:what parseText — "text" as UTF8-TEXT, "text"@tag as UTF8-LANGUAGE-TEXT with a LANGUAGE-BCP47-CODE language
package notation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseText reads a Go-quoted string, then either nothing or `@` and a
// language tag.
func parseText(tok string) (core.Record, error) {
	end := strings.LastIndex(tok, `"`)
	s, err := strconv.Unquote(tok[:end+1])
	if err != nil {
		return core.Record{}, fmt.Errorf("geul/notation: %s is not a quoted string", tok)
	}
	rest := tok[end+1:]
	switch {
	case rest == "":
		return core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s)), nil
	case strings.HasPrefix(rest, "@") && len(rest) > 1:
		lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(rest[1:]))
		return core.EncodeLanguageText(lang, []byte(s)), nil
	default:
		return core.Record{}, fmt.Errorf("geul/notation: %q after a quoted string", rest)
	}
}
