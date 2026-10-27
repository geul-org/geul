//ff:func feature=resolve type=resolver control=sequence
//ff:what fetchVersion — one listed version's bytes from a source, verified, as a Node
package resolve

import "context"

// fetchVersion fetches v's canonical address from e and verifies it as version
// v.Version of k.
func fetchVersion(ctx context.Context, e Entry, k Key, v VersionEntry) (*Node, error) {
	b, err := e.Source.Fetch(ctx, versionPath(k.DocKey(), v.Version))
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
