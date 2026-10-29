//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Archive.Compact — Compact every set (live records into fresh objects, one P and one H layer) and reopen it
package pack

import (
	"context"
	"errors"
)

// Compact rewrites each set of a local archive; staged changes must be committed first.
func (a *Archive) Compact(ctx context.Context) ([]BuildResult, error) {
	if a.dir == "" || a.Pending() {
		return nil, errors.New("geul/pack: compact needs a local archive without staged changes")
	}
	var out []BuildResult
	for _, s := range a.sets {
		if s.r == nil {
			continue
		}
		res, err := Compact(ctx, s.dir, a.o)
		if err == nil {
			s.r, err = Open(ctx, FileStore{Root: s.dir}, Options{Manifest: res.Manifest})
		}
		if err != nil {
			return out, err
		}
		out = append(out, res)
	}
	a.mu.Lock()
	a.idx = nil
	a.mu.Unlock()
	return out, nil
}
