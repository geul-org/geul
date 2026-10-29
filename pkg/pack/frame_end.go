//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what frameEnd — the length of the zstd frame at the start of b: header, blocks up to Last_Block, checksum
package pack

// frameEnd walks the blocks of the frame whose header is h, so a reader can
// insist on exactly one frame (§4.1) before handing it to the decoder.
func frameEnd(b []byte, h frameHeader) (int, error) {
	rest := b[h.size:]
	limit := min(h.window, maxBlock)
	for {
		var last bool
		var err error
		if rest, last, err = skipBlock(rest, limit); err != nil {
			return 0, err
		}
		if last {
			break
		}
	}
	if h.checksum && len(rest) < 4 {
		return 0, corrupt("zstd content checksum truncated")
	}
	if h.checksum {
		rest = rest[4:]
	}
	return len(b) - len(rest), nil
}
