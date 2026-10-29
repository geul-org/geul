//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.blockAt — a directory block from an index object, cached; zstd blocks must carry a checksum
package pack

import "context"

// blockAt reads, unframes and decodes the block at l, keeping up to r.limit
// blocks (the cache is dropped whole when full).
func (r *Reader) blockAt(ctx context.Context, l Loc) (block, error) {
	r.mu.Lock()
	b, ok := r.blocks[l]
	r.mu.Unlock()
	if ok {
		return b, nil
	}
	raw, err := r.raw(ctx, kindIndex, l)
	if err == nil {
		raw, err = r.unstore(ctx, l.Enc, raw, true)
	}
	if err == nil {
		b, err = decodeBlock(raw)
	}
	if err != nil {
		return block{}, err
	}
	r.mu.Lock()
	if len(r.blocks) >= r.limit {
		r.blocks = map[Loc]block{}
	}
	r.blocks[l] = b
	r.mu.Unlock()
	return b, nil
}
