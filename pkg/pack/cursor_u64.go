//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.u64 — a little-endian 64-bit integer
package pack

import "encoding/binary"

// u64 reads a little-endian uint64.
func (c *cursor) u64() uint64 {
	b := c.take(8)
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint64(b)
}
