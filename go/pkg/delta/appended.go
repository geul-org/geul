//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what appended — the diff's Edges whose IDs the base lacks, in the diff's order
package delta

import "github.com/geul-org/geul/go/pkg/core"

// appended returns the Edges of diff x still in left (written but not in the
// base), in x's order: the new statements a snapshot gets at its end.
func appended(x snapshot, left map[uint64]core.Record) []core.Record {
	var out []core.Record
	for i, id := range x.ids {
		if _, ok := left[id]; ok {
			out = append(out, x.body[i])
		}
	}
	return out
}
