//ff:func feature=resolve type=resolver control=sequence
//ff:what fetchFrom — one source: version list, chosen version, document bytes, verification
package resolve

import (
	"context"
	"fmt"
)

// fetchFrom resolves k against one source: read the version list, pick the
// version (latest for 0), fetch that version's canonical address and verify it.
// A version missing from the list is ErrNotFound for this source.
func fetchFrom(ctx context.Context, e Entry, k Key) (*Node, error) {
	d := k.DocKey()
	lb, err := e.Source.Fetch(ctx, listPath(d))
	if err != nil {
		return nil, err
	}
	list, err := parseVersionList(lb)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrVerify, e.Name, err)
	}
	v, ok := pickVersion(list, k.Version)
	if !ok {
		return nil, ErrNotFound
	}
	b, err := e.Source.Fetch(ctx, versionPath(d, v.Version))
	if err != nil {
		return nil, err
	}
	n, err := verifyDoc(b, k, v)
	if err != nil {
		return nil, err
	}
	n.Source = e.Name
	return n, nil
}
