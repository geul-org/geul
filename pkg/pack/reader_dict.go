//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.dict — a zstd dictionary's bytes, read once and checked against its DICTS SHA-256
package pack

import (
	"context"
	"crypto/sha256"
)

// dict returns dictionary id, fetching and caching it on first use.
func (r *Reader) dict(ctx context.Context, id uint32) ([]byte, error) {
	r.mu.Lock()
	d, ok := r.dicts[id]
	r.mu.Unlock()
	if ok {
		return d, nil
	}
	e, ok := findDict(r.m.dicts, id)
	if !ok {
		return nil, corrupt("zstd dictionary not in DICTS")
	}
	d, err := r.raw(ctx, kindData, e.loc)
	if err == nil && sha256.Sum256(d) != e.hash {
		err = corrupt("zstd dictionary SHA-256")
	}
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	r.dicts[id] = d
	r.mu.Unlock()
	return d, nil
}
