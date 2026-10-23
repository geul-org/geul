//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what declaredFields — a DECLARE's field list; false when it names a missing FIELD
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// declaredFields resolves d's FIELD Edge Ids in order. It is false when one is
// not a FIELD of the document.
func declaredFields(d core.Declare, fields map[uint64]*Field) ([]Field, bool) {
	out := make([]Field, 0, len(d.Fields))
	for _, id := range d.Fields {
		f := fields[id]
		if f == nil {
			return nil, false
		}
		out = append(out, *f)
	}
	return out, true
}
