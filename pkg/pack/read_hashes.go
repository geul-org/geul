//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readHashes — u64 count and 32-byte hashes (an inventory's deleted h/ list)
package pack

// readHashes reads the deleted hash list; failures are left in c.bad.
func readHashes(c *cursor) []string {
	n := c.u64()
	if c.bad || n > uint64(len(c.b))/32 {
		c.bad = true
		return nil
	}
	out := make([]string, n)
	for i := range out {
		out[i] = string(c.take(32))
	}
	return out
}
