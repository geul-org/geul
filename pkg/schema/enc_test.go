//ff:func feature=schema type=helper control=sequence
//ff:what enc — test fixture: a record's words
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
)

func enc(r core.Record) []core.Word {
	ws, _ := core.EncodeRecord(r)
	return ws
}
