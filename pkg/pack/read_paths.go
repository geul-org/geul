//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readPaths — an inventory path table: u32 count, u16 length and UTF-8 bytes each
package pack

// readPaths reads the path table; failures are left in c.bad.
func readPaths(c *cursor) []string {
	n := c.u32()
	if c.bad || uint64(n)*2 > uint64(len(c.b)) {
		c.bad = true
		return nil
	}
	out := make([]string, n)
	for i := range out {
		out[i] = string(c.take(uint64(c.u16())))
	}
	return out
}
