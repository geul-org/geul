//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what decodeWith — read the F_WITH count and TEXT REFs at *pos
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// decodeWith reads the `with` list at ws[*pos]: the count at Edge Id width, then
// that many TEXT REFs. A count larger than the remaining payload could hold is
// rejected before allocating.
func decodeWith(ws []core.Word, pos *int, w Wire) ([]TID, error) {
	n, err := TakeID(ws, pos, w.ID)
	if err != nil {
		return nil, err
	}
	if uint64(n) > uint64(len(ws)-*pos)/uint64(w.ID) {
		return nil, errors.New("geul/toulmin: F_WITH list longer than the payload")
	}
	out := make([]TID, n)
	for i := range out {
		out[i], _ = TakeID(ws, pos, w.ID)
	}
	return out, nil
}
