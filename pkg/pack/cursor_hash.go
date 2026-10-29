//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.hash — a 32-byte SHA-256 value
package pack

// hash reads 32 bytes; the zero hash when the cursor fails.
func (c *cursor) hash() [32]byte {
	var h [32]byte
	copy(h[:], c.take(32))
	return h
}
