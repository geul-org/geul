//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeID — read an Edge Id / REF field at *pos and advance past it
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// TakeID reads the width-word ID/REF field at ws[*pos] and advances *pos.
func TakeID(ws []core.Word, pos *int, width int) (TID, error) {
	v, err := core.DecodeID(ws[*pos:], width)
	if err != nil {
		return 0, err
	}
	*pos += width
	return TID(v), nil
}
