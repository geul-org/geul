//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.u16 — a little-endian 16-bit integer
package pack

import "encoding/binary"

// u16 reads a little-endian uint16.
func (c *cursor) u16() uint16 {
	b := c.take(2)
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint16(b)
}
