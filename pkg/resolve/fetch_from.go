//ff:func feature=resolve type=resolver control=sequence
//ff:what fetchFrom — one source: version list, pinned or newest version, verified document
package resolve

import "context"

// fetchFrom resolves k against one source: read the version list, pick the
// version (newest non-deprecated for 0) and fetch it. A version missing from the
// list is ErrNotFound for this source.
func fetchFrom(ctx context.Context, e Entry, k Key) (*Node, error) {
	list, err := sourceList(ctx, e, k.DocKey())
	if err != nil {
		return nil, err
	}
	v, ok := pickVersion(list, k.Version)
	if !ok {
		return nil, ErrNotFound
	}
	return fetchVersion(ctx, e, k, v)
}
