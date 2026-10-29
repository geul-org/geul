//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readInvEntries — u64 count and 32-byte inventory entries
package pack

// readInvEntries reads the entry list; failures are left in c.bad.
func readInvEntries(c *cursor) []invEntry {
	n := c.u64()
	if c.bad || n > uint64(len(c.b))/32 {
		c.bad = true
		return nil
	}
	out := make([]invEntry, n)
	for i := range out {
		out[i] = invEntry{kind: c.u8(), enc: c.u8(), flags: c.u16(), ref: c.u32(), doc: c.u64(), slot: c.u32(), n: c.u32(), off: c.u64()}
	}
	return out
}
