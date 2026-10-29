//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.Paths — every path this generation serves, sorted
package pack

import (
	"context"
	"sort"
)

// Paths lists all live paths (P, h/ and named).
func (r *Reader) Paths(ctx context.Context) ([]string, error) {
	es, err := r.entries(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(es))
	for p := range es {
		out = append(out, p)
	}
	sort.Strings(out)
	return out, nil
}
