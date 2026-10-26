//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what tokenize — split a line at top-level spaces, keeping quoted text, angle-bracket URIs and NAME(…) arguments whole
package notation

import (
	"errors"
	"strings"
)

// tokenize splits s at spaces and tabs outside quotes ("…" with backslash
// escapes), angle brackets and parentheses. An unclosed quote, bracket or
// parenthesis is an error.
func tokenize(s string) ([]string, error) {
	var out []string
	var cur strings.Builder
	sc := scanState{}
	for _, r := range s {
		if sc.top() && (r == ' ' || r == '\t') {
			out = flushToken(out, &cur)
			continue
		}
		cur.WriteRune(r)
		sc.step(r)
	}
	if !sc.top() {
		return nil, errors.New("geul/notation: unclosed quote, <…> or (…)")
	}
	return flushToken(out, &cur), nil
}
