//ff:func feature=pack type=resolver control=sequence
//ff:what OpenRemote — a read-only Archive over one pack set in any Store (an HTTP base, a mirror)
package pack

import "context"

// OpenRemote opens the set in s; WriteFile, Remove and Commit fail on it.
func OpenRemote(ctx context.Context, s Store, opt Options) (*Archive, error) {
	r, err := Open(ctx, s, opt)
	if err != nil {
		return nil, err
	}
	return &Archive{sets: []*packSet{{r: r}}, puts: map[string][]byte{}, dels: map[string]bool{}}, nil
}
