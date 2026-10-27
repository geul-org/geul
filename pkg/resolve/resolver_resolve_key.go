//ff:func feature=resolve type=resolver control=sequence
//ff:what Resolver.resolveKey — a pinned Key as is; a latest Key as the newest version, or the one in effect at AsOf
package resolve

import "context"

// resolveKey fetches k. A pinned version is fetched as is. A latest (version 0)
// Key is the newest version, or with AsOf set the version in effect then.
func (r *Resolver) resolveKey(ctx context.Context, k Key) (*Node, error) {
	if k.Version != 0 || r.AsOf.IsZero() {
		return r.Fetch(ctx, k)
	}
	return r.fetchWith(ctx, k, func(ctx context.Context, e Entry, k Key) (*Node, error) {
		return fetchAsOf(ctx, e, k, r.AsOf)
	})
}
