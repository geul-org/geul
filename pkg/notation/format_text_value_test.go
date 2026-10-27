//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatTextValue — UTF8-TEXT quoted, URI bracketed, other byte strings by name
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatTextValue(t *testing.T) {
	if formatTextValue(core.CodecUTF8TextEven, `a"b`) != `"a\"b"` || formatTextValue(codecURIEven, "https://x") != "<https://x>" || formatTextValue(core.CodecEDTFEven, "1971~") != `EDTF("1971~")` {
		t.Fatal("forms")
	}
}
