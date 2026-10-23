//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what fetchAsOf — one source: the version of a document in effect at an instant
package resolve

import (
	"context"
	"time"
)

// fetchAsOf finds the version in effect at at (OVERVIEW 계보 개정): the highest
// version whose EFFECTIVE-FROM is not after at. If that version's
// EFFECTIVE-UNTIL has passed, nothing is in effect and the result is
// ErrNotFound; older versions do not revive. Deprecation is ignored: a superseded
// version still applies to its own period.
func fetchAsOf(ctx context.Context, e Entry, k Key, at time.Time) (*Node, error) {
	list, err := sourceList(ctx, e, k.DocKey())
	if err != nil {
		return nil, err
	}
	for i := len(list) - 1; i >= 0; i-- {
		n, err := fetchVersion(ctx, e, k, list[i])
		if err != nil {
			return nil, err
		}
		st, err := effectiveState(n.Doc, at)
		if err != nil {
			return nil, err
		}
		if st == stateNotYet {
			continue
		}
		if st == stateExpired {
			return nil, ErrNotFound
		}
		return n, nil
	}
	return nil, ErrNotFound
}
