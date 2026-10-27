//ff:func feature=page type=helper control=selection
//ff:what textOf — the text of one UTF8-TEXT or UTF8-LANGUAGE-TEXT record (language text as "text"@lang); false for any other record
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// textOf decodes ws as one Value record and returns its text: a UTF8-TEXT as
// is, a UTF8-LANGUAGE-TEXT as the quoted text and its language ("Paris"@fr).
// ok is false for any other record.
func textOf(ws []core.Word) (string, bool) {
	r, _, err := core.DecodeRecord(ws)
	if err != nil || r.Kind != core.KindValue {
		return "", false
	}
	v, ok := inspect.ParseValue(r)
	switch {
	case !ok:
		return "", false
	case r.Codec&^1 == core.CodecUTF8TextEven:
		return v.Text, true
	case r.Codec&^1 == core.CodecUTF8LanguageTextEven:
		return strconv.Quote(v.Text) + "@" + v.Lang, true
	default:
		return "", false
	}
}
