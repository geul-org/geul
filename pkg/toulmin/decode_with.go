//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what decodeWith — read the F_WITH count Value and TEXT REFs at *pos
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// decodeWith reads the `with` list at ws[*pos]: the UNSIGNED-INTEGER count, then
// that many TEXT REFs. Each REF is at least 2 words, so a count larger than the
// remaining payload could hold is rejected before allocating.
func decodeWith(ws []core.Word, pos *int) ([]TID, error) {
	v, err := TakeValue(ws, pos)
	if err != nil {
		return nil, err
	}
	n, err := core.DecodeUintValue(v)
	if err != nil {
		return nil, err
	}
	if n > uint64(len(ws)-*pos)/2 {
		return nil, errors.New("geul/toulmin: F_WITH list longer than the payload")
	}
	out := make([]TID, n)
	for i := range out {
		if out[i], err = TakeID(ws, pos); err != nil {
			return nil, err
		}
	}
	return out, nil
}
