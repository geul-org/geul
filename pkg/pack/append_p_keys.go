//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what appendPKeys — C1·C2·C3 of ascending P keys, delta-coded against the previous key
package pack

// appendPKeys writes the three P key columns.
func appendPKeys(out []byte, keys []string) []byte {
	c1 := make([]uint64, len(keys))
	c2 := make([]uint64, len(keys))
	c3 := make([]uint64, len(keys))
	var pi uint32
	var pd uint64
	var ps uint32
	for j, k := range keys {
		i, d, s := splitPKey(k)
		c1[j], c2[j], c3[j] = pKeyDelta(j, pi, pd, ps, i, d, s)
		pi, pd, ps = i, d, s
	}
	return appendColumns(out, c1, c2, c3)
}
