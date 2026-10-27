//ff:func feature=page type=builder control=iteration dimension=1
//ff:what backrefMap — Edge Id → the Edge Ids of the records that reference it, in document order
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// backrefMap inverts every entry's references: for each referenced Edge Id,
// the Edge Ids of the records that reference it, in document order.
func backrefMap(body []inspect.Entry) map[uint64][]uint64 {
	back := map[uint64][]uint64{}
	for _, e := range body {
		for _, id := range e.Refs {
			back[id] = append(back[id], e.ID)
		}
	}
	return back
}
