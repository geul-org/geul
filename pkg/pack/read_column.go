//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readColumn — n varints of one column
package pack

// readColumn reads n varints; a failure is left in c.bad.
func readColumn(c *cursor, n int) []uint64 {
	v := make([]uint64, n)
	for i := range v {
		v[i] = c.uvarint()
	}
	return v
}
