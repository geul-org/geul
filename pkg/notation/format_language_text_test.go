//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatLanguageText — "text"@tag; a short language code and a broken record are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatLanguageText(t *testing.T) {
	short := core.Record{Kind: core.KindValue, Codec: 0x48, Payload: []core.Word{1}}
	if _, ok := formatLanguageText(core.EncodeLanguageText(short, []byte("x"))); ok {
		t.Fatal("short code: want false")
	}
	if _, ok := formatLanguageText(core.Record{Kind: core.KindValue, Codec: core.CodecUTF8LanguageTextEven}); ok {
		t.Fatal("broken: want false")
	}
}
