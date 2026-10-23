//ff:func feature=schema type=helper control=sequence
//ff:what recOf — test fixture: a rec around a Record
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
)

func recOf(r core.Record) rec {
	return rec{Index: 1, Word: 3, Total: len(enc(r)), R: r}
}
