//ff:func feature=notation type=helper control=sequence
//ff:what absent — test fixture: the Absent record (word 0x0000)
package notation

import "github.com/geul-org/geul/go/pkg/core"

// absent is the record word 0x0000 decodes to.
func absent() core.Record {
	r, _, _ := core.DecodeRecord([]core.Word{0})
	return r
}
