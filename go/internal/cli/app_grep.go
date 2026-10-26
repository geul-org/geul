//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.grep — print `document: line` for each matching line of the given documents or of the whole data folder
package cli

import (
	"regexp"
)

// grep matches pattern against every HEAD and BODY line (linesOf) of each
// target (grepTargets).
func (a *app) grep(pattern string, args []string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	targets, err := a.grepTargets(args)
	if err != nil {
		return err
	}
	for _, t := range targets {
		b, err := a.readDoc(t)
		if err != nil {
			return err
		}
		ls, err := a.linesOf(b)
		if err != nil {
			return err
		}
		a.printMatches(t, re, append(ls.Head, ls.Body...))
	}
	return nil
}
