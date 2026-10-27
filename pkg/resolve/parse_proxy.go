//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what parseProxy — a GEULPROXY list split on `,` and `|`, recording which separator follows each item
package resolve

import "strings"

// parseProxy splits s into entries. `,` and `|` both separate items; an item
// followed by `|` falls through on any error (GEUL-ORG GEULPROXY).
func parseProxy(s string) ([]Entry, error) {
	var out []Entry
	start := 0
	for i := 0; i <= len(s); i++ {
		if i < len(s) && s[i] != ',' && s[i] != '|' {
			continue
		}
		e, err := parseEntry(strings.TrimSpace(s[start:i]))
		if err != nil {
			return nil, err
		}
		e.Fallthrough = i < len(s) && s[i] == '|'
		out = append(out, e)
		start = i + 1
	}
	return out, nil
}
