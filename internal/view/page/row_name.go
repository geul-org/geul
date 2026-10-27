//ff:func feature=page type=helper control=sequence
//ff:what rowName — the spec name of a record, or its kind and codec when it has none
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/inspect"
)

// rowName returns e's spec name, or `<KIND> 0x<codec>` for a codec no spec
// name is known for.
func rowName(e inspect.Entry) string {
	if e.Name != "" {
		return e.Name
	}
	return fmt.Sprintf("%s 0x%02X", e.Record.Kind, e.Record.Codec)
}
