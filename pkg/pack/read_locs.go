//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readLocs — n locations from the C4–C7 columns
package pack

// readLocs reads the location columns of an n-entry block.
func readLocs(c *cursor, n int, leaf bool) ([]Loc, error) {
	c4, c5, c6 := readColumn(c, n), readColumn(c, n), readColumn(c, n)
	c7 := c.take(uint64(n))
	if c.bad {
		return nil, corrupt("location columns")
	}
	locs := make([]Loc, n)
	for j := range locs {
		l, ok := locStep(locs, j, c4[j], c5[j], c6[j], c7[j], leaf)
		if !ok {
			return nil, corrupt("location entry")
		}
		locs[j] = l
	}
	return locs, nil
}
