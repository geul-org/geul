//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseDicts — the DICTS section: u32 count, 52-byte entries (ID, object, offset, length, SHA-256)
package pack

// parseDicts decodes DICTS.
func parseDicts(b []byte) ([]dictEntry, error) {
	c := &cursor{b: b}
	n := c.u32()
	if uint64(n)*52 != uint64(len(c.b)) {
		return nil, corrupt("DICTS length")
	}
	out := make([]dictEntry, n)
	for i := range out {
		id, obj, off, ln := c.u32(), c.u32(), c.u64(), c.u32()
		out[i] = dictEntry{id: id, loc: Loc{Obj: obj, Off: off, Len: uint64(ln)}, hash: c.hash()}
	}
	return out, nil
}
