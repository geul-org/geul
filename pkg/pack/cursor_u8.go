//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.u8 — one byte
package pack

// u8 reads one byte.
func (c *cursor) u8() byte {
	b := c.take(1)
	if b == nil {
		return 0
	}
	return b[0]
}
