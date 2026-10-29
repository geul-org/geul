//ff:func feature=pack type=decoder control=selection
//ff:what zstdBlock — decode one zstd block (raw or RLE); a compressed block is ErrUnsupported
package pack

import (
	"bytes"
	"fmt"
)

// zstdBlock appends one block's content to out and returns the bytes after
// it and whether it was the last block.
func zstdBlock(out, b []byte, limit uint64) ([]byte, []byte, bool, error) {
	c := &cursor{b: b}
	hb := c.take(3)
	if hb == nil || uint64(hb[0]>>3)|uint64(hb[1])<<5|uint64(hb[2])<<13 > limit {
		return nil, nil, false, corrupt("zstd block header")
	}
	h := uint32(hb[0]) | uint32(hb[1])<<8 | uint32(hb[2])<<16
	last, n := h&1 == 1, uint64(h>>3)
	switch (h >> 1) & 3 {
	case 0:
		body := c.take(n)
		if c.bad {
			return nil, nil, false, corrupt("zstd raw block truncated")
		}
		return append(out, body...), c.b, last, nil
	case 1:
		v := c.u8()
		if c.bad {
			return nil, nil, false, corrupt("zstd RLE block truncated")
		}
		return append(out, bytes.Repeat([]byte{v}, int(n))...), c.b, last, nil
	case 2:
		return nil, nil, false, fmt.Errorf("%w: compressed zstd block (the built-in codec reads raw and RLE blocks only)", ErrUnsupported)
	default:
		return nil, nil, false, corrupt("reserved zstd block type")
	}
}
