//ff:func feature=schema type=helper control=iteration dimension=1
//ff:what rulesOf — test fixture: the rule IDs of Facts, joined
package schema

import (
	"strings"
)

func rulesOf(fs []Fact) string {
	var out []string
	for _, f := range fs {
		out = append(out, f.Rule)
	}
	return strings.Join(out, ",")
}
