//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeSpan — read a SPAN field at *pos and advance past it
package toulmin

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// TakeSpan reads the SPAN field (offset, length at width words each) at ws[*pos]
// and advances *pos.
func TakeSpan(ws []core.Word, pos *int, width int) (Span, error) {
	s, err := DecodeSpan(ws[*pos:], width)
	if err != nil {
		return Span{}, err
	}
	*pos += 2 * width
	return s, nil
}
