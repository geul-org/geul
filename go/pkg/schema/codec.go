//ff:type feature=schema type=model
//ff:what Codec — a layer-2 codec specification: which records it covers and how to check them
package schema

import "github.com/geul-org/geul/go/pkg/core"

// Codec is a layer-2 check (README 층): a codec specification that knows its
// records' payload layout. Check runs only on records Match accepts.
type Codec interface {
	Match(r core.Record) bool
	Check(r core.Record, at Where) []Fact
}
