//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseObjects — the OBJECTS section: u32 count, 48-byte entries (kind, reserved, size, SHA-256)
package pack

// parseObjects decodes OBJECTS; kinds other than 1 and 2 are ErrCorrupt.
func parseObjects(b []byte) ([]object, error) {
	c := &cursor{b: b}
	n := c.u32()
	if uint64(n)*48 != uint64(len(c.b)) {
		return nil, corrupt("OBJECTS length")
	}
	out := make([]object, n)
	for i := range out {
		kind, res := c.u8(), c.take(7)
		out[i] = object{kind: kind, size: c.u64(), hash: c.hash()}
		if kind != kindData && kind != kindIndex || !zeros(res) || out[i].size < headerSize {
			return nil, corrupt("OBJECTS entry")
		}
	}
	return out, nil
}
