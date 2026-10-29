//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.fetch — a path's bytes from the first set that serves it; false when none does
package pack

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// fetch searches the sets for name (staged changes not included).
func (a *Archive) fetch(name string) ([]byte, bool, error) {
	for _, s := range a.order(name) {
		b, err := s.r.Fetch(context.Background(), name)
		if err == nil {
			return b, true, nil
		}
		if !errors.Is(err, resolve.ErrNotFound) {
			return nil, false, err
		}
	}
	return nil, false, nil
}
