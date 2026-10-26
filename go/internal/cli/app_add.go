//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.add — parse each line as an Edge with labels mapped to the IDs assigned in order, append them, save, print the IDs
package cli

import (
	"errors"
	"fmt"

	"github.com/geul-org/geul/go/pkg/notation"
)

// add gives line i the Edge Id next+i (nextFree) and maps each line's
// #label to its ID (labelIDs), so lines can refer to each other. Every line
// must read (notation.ParseEdge) before anything is written; the draft is
// saved only when it keeps the spec (saveDraft). Each line's ID is printed
// as #n.
func (a *app) add(path string, lines []string) error {
	if len(lines) == 0 || lines[0] == "" {
		return errors.New("add: no Edge given")
	}
	d, err := readDraft(path)
	if err != nil {
		return err
	}
	next := a.nextFree(d)
	labels := labelIDs(lines, next)
	ctx := a.context()
	ctx.Local = func(n uint64) uint64 {
		if id, ok := labels[n]; ok {
			return id
		}
		return n
	}
	for i, line := range lines {
		e, err := notation.ParseEdge(line, ctx)
		if err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}
		d.body = append(d.body, edgeRecord(next+uint64(i), e))
	}
	if err := saveDraft(path, d); err != nil {
		return err
	}
	for i := range lines {
		fmt.Fprintf(a.out, "#%d\n", next+uint64(i))
	}
	return nil
}
