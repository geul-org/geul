//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Resolver.fetchWith — try each source in order with the GEULPROXY and GEULPRIVATE rules
package resolve

import (
	"context"
	"errors"
	"fmt"
)

// fetchWith runs f against the sources in order. geul.org and registry are
// skipped for a private document; `off` stops. A source that lacks the document
// passes to the next one, as does any error of a source followed by `|`. A
// verification failure always stops.
func (r *Resolver) fetchWith(ctx context.Context, k Key, f fetchFunc) (*Node, error) {
	d := k.DocKey()
	private := matchAny(r.Config.Private, d)
	for _, e := range r.Config.Proxy {
		if e.Off {
			break
		}
		if e.Public && private {
			continue
		}
		n, err := f(ctx, e, k)
		if err == nil {
			return n, nil
		}
		if errors.Is(err, ErrVerify) || (!errors.Is(err, ErrNotFound) && !e.Fallthrough) {
			return nil, fmt.Errorf("geul/resolve: %s/%s from %s: %w", d.Issuer, d.Doc, e.Name, err)
		}
	}
	return nil, fmt.Errorf("geul/resolve: %s/%s: %w", d.Issuer, d.Doc, ErrNotFound)
}
