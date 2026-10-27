//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestTokenize — spaces split only at the top level: quotes with escapes, URIs and calls stay whole; unclosed ones fail
package notation

import (
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {
	got, err := tokenize(`#1  LABEL Q42	"a \" b"@en <https://x y> QUANTITY(1, _, (2), "a)b")`)
	want := []string{"#1", "LABEL", "Q42", `"a \" b"@en`, "<https://x y>", `QUANTITY(1, _, (2), "a)b")`}
	if err != nil || !slices.Equal(got, want) {
		t.Fatalf("%q %v", got, err)
	}
	for _, s := range []string{`"open`, "<open", "CALL(open"} {
		if _, err := tokenize(s); err == nil {
			t.Fatalf("%s: want error", s)
		}
	}
}
