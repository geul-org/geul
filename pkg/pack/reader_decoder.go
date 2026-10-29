//ff:func feature=pack type=decoder control=sequence
//ff:what Reader.decoder — the zstd decoder for a frame's Dictionary_ID: the shared one for 0, else one per DICTS dictionary, cached
package pack

import (
	"context"

	"github.com/klauspost/compress/zstd"
)

// decoder returns the decoder for dictionary id, loading the dictionary
// (SHA-256 checked) on first use.
func (r *Reader) decoder(ctx context.Context, id uint32) (*zstd.Decoder, error) {
	if id == 0 {
		return plainDecoder(), nil
	}
	r.mu.Lock()
	d, ok := r.decs[id]
	r.mu.Unlock()
	if ok {
		return d, nil
	}
	b, err := r.dict(ctx, id)
	if err == nil {
		d, err = newDictDecoder(id, b)
	}
	if err != nil {
		return nil, err
	}
	r.mu.Lock()
	r.decs[id] = d
	r.mu.Unlock()
	return d, nil
}
