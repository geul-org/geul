//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.Issuers — Issuers holding any live document path or named path under them, in resolve.Dir order
package pack

import (
	"context"
	"sort"
)

// Issuers lists the canonical Issuer IDs this generation has paths under
// (`<I>/…`): live P entries or named paths.
func (r *Reader) Issuers(ctx context.Context) ([]string, error) {
	var out []string
	for i, id := range r.m.issuers {
		s, err := r.live(ctx, r.m.pLayers, pKey(uint32(i), 0, 0), pKey(uint32(i)+1, 0, 0))
		if err != nil {
			return nil, err
		}
		if len(s) > 0 || r.hasNamedUnder(id+"/") {
			out = append(out, id)
		}
	}
	sort.Slice(out, func(a, b int) bool { return issuerLess(out[a], out[b]) })
	return out, nil
}
