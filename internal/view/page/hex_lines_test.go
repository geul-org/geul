//ff:func feature=page type=helper control=sequence
//ff:what TestHexLines — one line per word with the first named; no words give no lines; a long run is cut
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHexLines(t *testing.T) {
	got := hexLines([]core.Word{1, 0xABCD}, "flags")
	if len(got) != 2 || got[0].Word != "0001" || got[0].Name != "flags" || got[1].Word != "ABCD" || got[1].Name != "" {
		t.Fatalf("short: %+v", got)
	}
	if got := hexLines(nil, "x"); len(got) != 0 {
		t.Fatalf("empty: %+v", got)
	}
	got = hexLines(make([]core.Word, hexLimit+5), "text")
	if len(got) != hexLimit+1 || got[hexLimit].Word != "…" || got[hexLimit].Note != "+5 WORDs" || got[0].Name != "text" {
		t.Fatalf("long: %+v", got[hexLimit])
	}
}
