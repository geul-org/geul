//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what readPKeys — n P keys from the C1·C2·C3 columns, rejecting overflow and non-increasing keys
package pack

// readPKeys reads and rebuilds the P key columns.
func readPKeys(c *cursor, n int) ([]string, error) {
	c1, c2, c3 := readColumn(c, n), readColumn(c, n), readColumn(c, n)
	if c.bad {
		return nil, corrupt("P key columns")
	}
	keys := make([]string, n)
	var pi uint32
	var pd uint64
	var ps uint32
	for j := range keys {
		var ok bool
		pi, pd, ps, ok = pKeyStep(j, pi, pd, ps, c1[j], c2[j], c3[j])
		if !ok {
			return nil, corrupt("P key out of range or not increasing")
		}
		keys[j] = pKey(pi, pd, ps)
	}
	return keys, nil
}
