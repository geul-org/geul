//ff:func feature=schema type=rule control=sequence
//ff:what fakeCodec.Check — test fixture: one X1 Fact at the record
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
)

func (f fakeCodec) Check(_ core.Record, at Where) []Fact {
	return []Fact{{Rule: "X1", Severity: Error, Where: at}}
}
