//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what encodeWith — append the F_WITH field: a count at Edge Id width then one TEXT REF per term
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// encodeWith appends the `with` list: the item count, written at the Edge Id
// width (a list of REFs can never outnumber the document's Edge Ids, so that
// width always holds it), then one TEXT REF per term.
func encodeWith(out []core.Word, with []TID, w Wire) ([]core.Word, error) {
	out, err := PutID(out, TID(len(with)), w.ID)
	for _, t := range with {
		if err != nil {
			return nil, err
		}
		out, err = PutID(out, t, w.ID)
	}
	return out, err
}
