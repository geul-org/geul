//ff:func feature=toulmin type=encoder control=sequence
//ff:what PutID — append an Edge Id / REF field at the declared width
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// PutID appends v as a width-word ID/REF field to out. A value that does not fit
// the declared width is an error.
func PutID(out []core.Word, v TID, width int) ([]core.Word, error) {
	ws, err := core.EncodeID(uint64(v), width)
	if err != nil {
		return nil, err
	}
	return append(out, ws...), nil
}
