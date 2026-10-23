//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeID — read an embedded LOCAL Reference (Edge Id / REF) at *pos and advance past it
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// TakeID decodes the embedded record at ws[*pos], which must be a LOCAL
// Reference with Quantifier Identified, and advances *pos by its Length.
func TakeID(ws []core.Word, pos *int) (TID, error) {
	r, n, err := core.DecodeRecord(ws[*pos:])
	if err != nil {
		return 0, err
	}
	v, err := core.DecodeLocalRef(r)
	if err != nil {
		return 0, err
	}
	*pos += n
	return TID(v), nil
}
