//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what parsePatterns — a comma-separated pattern list; empty items are skipped
package resolve

import "strings"

// parsePatterns reads a GEULPRIVATE-style list.
func parsePatterns(s string) ([]Pattern, error) {
	var out []Pattern
	for _, item := range strings.Split(s, ",") {
		if item = strings.TrimSpace(item); item == "" {
			continue
		}
		p, err := parsePattern(item)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}
