//ff:func feature=cli type=helper control=sequence
//ff:what changeLines — a delta.Change with its old and new records written as lines of their documents
package cli

import (
	"github.com/geul-org/geul/go/pkg/delta"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// changeLines writes c.Old as a line of base and c.New as a line of next
// (lineOf), so each is laid out with its own document's specs.
func changeLines(c delta.Change, base, next inspect.Document) changeLine {
	l := changeLine{Op: string(c.Op), Head: c.Head, ID: c.ID}
	if c.Op != '+' {
		l.Old = lineOf(base, c.Old)
	}
	if c.Op != '-' {
		l.New = lineOf(next, c.New)
	}
	return l
}
