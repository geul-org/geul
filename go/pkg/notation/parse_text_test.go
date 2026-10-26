//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseText — UTF8-TEXT, language text with a tag; a broken quote, a bare @ and junk after the quote fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseText(t *testing.T) {
	if r, err := parseText(`"a\"b"`); err != nil || r.Codec&^1 != core.CodecUTF8TextEven {
		t.Fatalf("%+v %v", r, err)
	}
	r, err := parseText(`"Adams"@en`)
	if s, ok := formatLanguageText(r); err != nil || !ok || s != `"Adams"@en` {
		t.Fatalf("%s %v", s, err)
	}
	for _, s := range []string{`"a`, `"a"@`, `"a"x`} {
		if _, err := parseText(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
