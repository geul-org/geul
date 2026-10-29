//ff:func feature=pack type=writer control=sequence
//ff:what Archive.commitSet — write one set's batch and reopen the set at its new generation
package pack

import (
	"context"
	"errors"
)

// commitSet writes b to s and reopens s.
func (a *Archive) commitSet(ctx context.Context, s *packSet, b *Batch) (BuildResult, error) {
	if a.dir == "" {
		return BuildResult{}, errors.New("geul/pack: read-only pack archive")
	}
	res, err := Write(ctx, s.dir, *b, a.o)
	if err != nil || !res.Changed {
		return res, err
	}
	s.r, err = Open(ctx, FileStore{Root: s.dir}, Options{Manifest: res.Manifest})
	return res, err
}
