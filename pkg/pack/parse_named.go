//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseNamed — the NAMED section: u32 count, path, encoding, object, offset, length; strictly ascending paths
package pack

// parseNamed decodes NAMED; paths must ascend strictly and encodings be 0 or 1.
func parseNamed(b []byte) ([]namedEntry, error) {
	c := &cursor{b: b}
	n := c.u32()
	if c.bad || uint64(n)*20 > uint64(len(c.b)) {
		return nil, corrupt("NAMED length")
	}
	out := make([]namedEntry, n)
	for i := range out {
		p := string(c.take(uint64(c.u16())))
		enc, res := c.u8(), c.u8()
		out[i] = namedEntry{path: p, loc: Loc{Obj: c.u32(), Off: c.u64(), Len: uint64(c.u32()), Enc: enc}}
		if c.bad || enc > encZstd || res != 0 || i > 0 && out[i-1].path >= p {
			return nil, corrupt("NAMED entry")
		}
	}
	return out, c.end("NAMED")
}
