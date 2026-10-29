//ff:func feature=notation type=decoder control=sequence
//ff:what namedLanguageToken — a "text"@@0xNNNN token with the registry tag in place of the LANGUAGE-SHORT-CODE hex
package notation

import "strings"

// namedLanguageToken rewrites the language of a language text token
// (formatLanguageText) written as `"…"@@0xNNNN` to `"…"@@tag` when langs (the
// GEUL Language Registry, inspect.Entry.Langs) names that code. Every other
// token, and a code langs does not have, is returned as is.
func namedLanguageToken(tok string, langs map[string]string) string {
	i := strings.LastIndex(tok, `"@@`)
	if i < 0 || !strings.HasPrefix(tok, `"`) {
		return tok
	}
	tag, ok := langs[tok[i+3:]]
	if !ok {
		return tok
	}
	return tok[:i+3] + tag
}
