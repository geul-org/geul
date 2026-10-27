//ff:func feature=resolve type=resolver control=sequence
//ff:what Resolver.Fetch — one Key, pinned or newest, through the source list
package resolve

import "context"

// Fetch resolves one Key: the pinned version, or the newest one for version 0
// (GEUL-ORG 위치 해석 순서).
func (r *Resolver) Fetch(ctx context.Context, k Key) (*Node, error) {
	return r.fetchWith(ctx, k, fetchFrom)
}
