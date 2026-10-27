//ff:type feature=resolve type=resolver
//ff:what fetchFunc — resolves a Key against one source (newest/pinned or as of an instant)
package resolve

import "context"

// fetchFunc resolves a Key against one source.
type fetchFunc func(ctx context.Context, e Entry, k Key) (*Node, error)
