//ff:func feature=page type=builder control=sequence
//ff:what inspectEntry — test fixture: an Entry of a record with its words and codec spec
package page

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// inspectEntry builds the Entry of r carrying spec.
func inspectEntry(r core.Record, spec []codecspec.Field) inspect.Entry {
	ws, _ := core.EncodeRecord(r)
	return inspect.Entry{Record: r, Words: ws, Spec: spec}
}
