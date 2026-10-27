//ff:type feature=resolve type=resolver
//ff:what Source — fetches a path under `.well-known/geul/` from one place
package resolve

import "context"

// Source fetches one path relative to `.well-known/geul/` (a document, a
// version list ending in `/`). A source that does not have it returns an error
// wrapping ErrNotFound; any other error means the source is broken.
type Source interface {
	Fetch(ctx context.Context, path string) ([]byte, error)
}
