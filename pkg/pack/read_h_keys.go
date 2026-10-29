//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readHKeys — n 32-byte H keys
package pack

// readHKeys reads the H key column; a short read is left in c.bad.
func readHKeys(c *cursor, n int) []string {
	keys := make([]string, n)
	for j := range keys {
		keys[j] = string(c.take(32))
	}
	return keys
}
