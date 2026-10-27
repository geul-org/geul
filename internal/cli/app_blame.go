//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.blame — read the versions up to the asked one and print delta.Blame's lines
package cli

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/delta"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// blame reads every live version of addr's document up to the asked one
// (blameVersions), blames them (delta.Blame) and prints `added changed line`.
func (a *app) blame(addr string) error {
	vs, err := a.blameVersions(addr)
	if err != nil {
		return err
	}
	ls, err := delta.Blame(vs)
	if err != nil {
		return err
	}
	doc, _ := inspect.DecodeWith(vs[len(vs)-1].Doc, a.spec())
	out := []blameLine{}
	for _, l := range ls {
		out = append(out, blameLine{Added: l.Added, Changed: l.Changed, Line: lineOf(doc, l.Record)})
	}
	if done, err := a.json(out); done {
		return err
	}
	for _, l := range out {
		fmt.Fprintf(a.out, "%d\t%d\t%s\n", l.Added, l.Changed, l.Line)
	}
	return nil
}
