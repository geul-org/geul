//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what splitArgs — the arguments of NAME(…): split at top-level commas, each trimmed
package notation

import (
	"strings"
)

// splitArgs splits the text between a call's parentheses at commas outside
// quotes, URIs and nested calls. Empty text has no arguments.
func splitArgs(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var out []string
	var cur strings.Builder
	sc := scanState{}
	for _, r := range s {
		if sc.top() && r == ',' {
			out = append(out, strings.TrimSpace(cur.String()))
			cur.Reset()
			continue
		}
		cur.WriteRune(r)
		sc.step(r)
	}
	return append(out, strings.TrimSpace(cur.String()))
}
