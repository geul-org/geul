//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.printMatches — the lines that match, prefixed with their document
package cli

import (
	"fmt"
	"regexp"
)

// printMatches writes `name: line` for each line re matches.
func (a *app) printMatches(name string, re *regexp.Regexp, lines []string) {
	for _, l := range lines {
		if re.MatchString(l) {
			fmt.Fprintf(a.out, "%s: %s\n", name, l)
		}
	}
}
