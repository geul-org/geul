//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.index — the sorted live paths of every set with staged changes applied (built once, then kept up to date)
package pack

import (
	"context"
	"slices"
)

// index returns the path listing; the caller holds a.mu.
func (a *Archive) index() ([]string, error) {
	if a.idx != nil {
		return a.idx, nil
	}
	seen := map[string]bool{}
	for _, s := range a.sets {
		if s.r == nil {
			continue
		}
		ps, err := s.r.Paths(context.Background())
		if err != nil {
			return nil, err
		}
		for _, p := range ps {
			seen[p] = true
		}
	}
	for p := range a.puts {
		seen[p] = true
	}
	for p := range a.dels {
		delete(seen, p)
	}
	a.idx = make([]string, 0, len(seen))
	for p := range seen {
		a.idx = append(a.idx, p)
	}
	slices.Sort(a.idx)
	return a.idx, nil
}
