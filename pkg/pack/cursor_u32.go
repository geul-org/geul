//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.u32 — a little-endian 32-bit integer
package pack

import "encoding/binary"

// u32 reads a little-endian uint32.
func (c *cursor) u32() uint32 {
	b := c.take(4)
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}
