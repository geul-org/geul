//ff:func feature=schema type=rule control=sequence
//ff:what fakeCodec.Match — test fixture: matches records of its codec
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
)

func (f fakeCodec) Match(r core.Record) bool {
	return r.Codec == f.codec
}
