//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestNamedLanguageToken — a "…"@@hex token the registry has gets its tag; an unknown code, a BCP 47 token, a non-text token and no registry stay
package notation

import "testing"

func TestNamedLanguageToken(t *testing.T) {
	langs := map[string]string{"0x0784": "en", "0x0038": "ach"}
	cases := map[string]string{
		`"Douglas Adams"@@0x0784`: `"Douglas Adams"@@en`,
		`"a ""@@ b"@@0x0038`:      `"a ""@@ b"@@ach`,
		`"x"@@0x7000`:             `"x"@@0x7000`,
		`"x"@en`:                  `"x"@en`,
		`#3`:                      `#3`,
		`0x0038`:                  `0x0038`,
	}
	for in, want := range cases {
		if got := namedLanguageToken(in, langs); got != want {
			t.Errorf("%s: got %s, want %s", in, got, want)
		}
	}
	if got := namedLanguageToken(`"x"@@0x0784`, nil); got != `"x"@@0x0784` {
		t.Errorf("no registry: %s", got)
	}
}
