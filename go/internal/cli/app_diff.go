//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.diff — print the changes between two documents (delta.Changes) as notation lines
package cli

import (
	"github.com/geul-org/geul/go/pkg/delta"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// diff resolves the pair (diffPair), compares them (delta.Changes) and prints
// each change with its records' lines (changeLines), or them as JSON.
func (a *app) diff(args []string) error {
	base, next, err := a.diffPair(args)
	if err != nil {
		return err
	}
	cs, err := delta.Changes(base, next)
	if err != nil {
		return err
	}
	bd, _ := inspect.DecodeWith(base, a.spec())
	nd, _ := inspect.DecodeWith(next, a.spec())
	var out []changeLine
	for _, c := range cs {
		out = append(out, changeLines(c, bd, nd))
	}
	if done, err := a.json(out); done {
		return err
	}
	for _, c := range out {
		a.printChange(c)
	}
	return nil
}
