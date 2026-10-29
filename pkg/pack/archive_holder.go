//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.holder — the set that already serves a path (so an overwrite stays where the path is), nil when none
package pack

import "context"

// holder finds the first set in read order that has name.
func (a *Archive) holder(ctx context.Context, name string) (*packSet, error) {
	for _, s := range a.order(name) {
		ok, err := s.r.Has(ctx, name)
		if err != nil || ok {
			return s, err
		}
	}
	return nil, nil
}
