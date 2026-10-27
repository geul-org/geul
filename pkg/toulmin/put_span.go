//ff:func feature=toulmin type=encoder control=sequence
//ff:what PutSpan — append a SPAN field at the declared SPAN width
package toulmin

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// PutSpan appends s as its offset and length fields at width words each.
func PutSpan(out []core.Word, s Span, width int) ([]core.Word, error) {
	ws, err := EncodeSpan(s, width)
	if err != nil {
		return nil, err
	}
	return append(out, ws...), nil
}
