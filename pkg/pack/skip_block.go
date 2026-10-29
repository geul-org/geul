//ff:func feature=pack type=decoder control=selection
//ff:what skipBlock — step over one zstd block (raw, RLE or compressed) by its header; reserved, oversize and truncated blocks are ErrCorrupt
package pack

// skipBlock returns the bytes after one block and whether it was the last.
// Blocks hold at most limit bytes (min(window, 128 KiB)); a compressed block's
// Block_Size is its stored size, an RLE block stores one byte.
func skipBlock(b []byte, limit uint64) ([]byte, bool, error) {
	c := &cursor{b: b}
	hb := c.take(3)
	if hb == nil {
		return nil, false, corrupt("zstd block header")
	}
	h := uint32(hb[0]) | uint32(hb[1])<<8 | uint32(hb[2])<<16
	last, n := h&1 == 1, uint64(h>>3)
	if n > limit {
		return nil, false, corrupt("zstd block larger than the window")
	}
	switch (h >> 1) & 3 {
	case 0, 2:
		c.take(n)
	case 1:
		c.u8()
	default:
		return nil, false, corrupt("reserved zstd block type")
	}
	if c.bad {
		return nil, false, corrupt("zstd block truncated")
	}
	return c.b, last, nil
}
