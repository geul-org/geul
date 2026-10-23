//ff:func feature=resolve type=resolver control=sequence
//ff:what sourceList — a document's version list from one source
package resolve

import (
	"context"
	"fmt"
)

// sourceList fetches and parses d's version list from e. A list that does not
// parse is a verification failure.
func sourceList(ctx context.Context, e Entry, d DocKey) ([]VersionEntry, error) {
	lb, err := e.Source.Fetch(ctx, listPath(d))
	if err != nil {
		return nil, err
	}
	list, err := parseVersionList(lb)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrVerify, e.Name, err)
	}
	return list, nil
}
