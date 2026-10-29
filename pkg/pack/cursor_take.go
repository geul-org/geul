//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.take — the next n bytes, nil and a sticky failure when fewer remain
package pack

// take returns the next n bytes or nil when the cursor has failed or fewer remain.
func (c *cursor) take(n uint64) []byte {
	if c.bad || n > uint64(len(c.b)) {
		c.bad = true
		return nil
	}
	out := c.b[:n:n]
	c.b = c.b[n:]
	return out
}
