//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what appendLocs — the C4–C7 location columns (object, offset code, length, encoding)
package pack

// appendLocs writes the location columns; a deletion marker is C4 = 0, C5 = 1, C6 = 0.
func appendLocs(out []byte, locs []Loc) []byte {
	c4 := make([]uint64, len(locs))
	c5 := make([]uint64, len(locs))
	c6 := make([]uint64, len(locs))
	c7 := make([]byte, len(locs))
	for j, l := range locs {
		c4[j], c5[j], c6[j], c7[j] = uint64(l.Obj), offCode(locs, j), l.Len, l.Enc
	}
	out = appendColumns(out, c4, c5, c6)
	return append(out, c7...)
}
