//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what Archive.Verify — Reader.Verify of every set (objects hashed, inventories replayed against the index)
package pack

import (
	"context"
	"fmt"
)

// Verify checks every set; the first failure names its set.
func (a *Archive) Verify(ctx context.Context) error {
	for _, s := range a.sets {
		if s.r == nil {
			continue
		}
		if err := s.r.Verify(ctx); err != nil {
			return fmt.Errorf("set %q: %w", s.name, err)
		}
	}
	return nil
}
