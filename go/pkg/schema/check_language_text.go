//ff:func feature=schema type=rule control=sequence
//ff:what checkLanguageText — V7: a UTF8-LANGUAGE-TEXT starts with a whole language record that is itself valid, then UTF-8 text with odd padding
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkLanguageText applies V7 to a UTF8-LANGUAGE-TEXT Value: its Payload
// starts with one whole LANGUAGE-SHORT-CODE or LANGUAGE-BCP47-CODE record,
// which must pass its own Value rules (checkValue), and the words after it
// are the text, which follows V4 (odd padding) and V5 (UTF-8, no BOM) as
// UTF8-TEXT does.
func checkLanguageText(r rec) []Fact {
	lang, n, err := core.DecodeRecord(r.R.Payload)
	if err != nil || lang.Kind != core.KindValue || !isLanguage(lang.Codec) {
		return []Fact{factAt("V7", Error, r, "Payload", "LANGUAGE-SHORT-CODE 또는 LANGUAGE-BCP47-CODE 레코드로 시작", fmt.Sprintf("%v", err))}
	}
	if fs := checkValue(rec{Index: r.Index, Word: r.Word, Total: n, R: lang}); len(fs) > 0 {
		return fs
	}
	text := r
	text.R = core.Record{Kind: core.KindValue, Codec: core.CodecUTF8TextEven | r.R.Codec&1, Payload: r.R.Payload[n:]}
	return checkBytes(text)
}
