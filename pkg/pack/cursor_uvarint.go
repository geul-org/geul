//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.uvarint — an unsigned LEB128 value in its shortest form (padding zero bytes rejected)
package pack

import "encoding/binary"

// uvarint reads an unsigned LEB128 value below 2^64 in at most 10 bytes and
// only in its shortest form: a final 0x00 continuation byte is malformed.
func (c *cursor) uvarint() uint64 {
	if c.bad {
		return 0
	}
	v, n := binary.Uvarint(c.b)
	if n <= 0 || n > 1 && c.b[n-1] == 0 {
		c.bad = true
		return 0
	}
	c.b = c.b[n:]
	return v
}
