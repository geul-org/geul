//ff:func feature=pack type=encoder control=selection
//ff:what offCode — the C5 value of entry j: 1 for a deletion marker, 0 when contiguous, offset + 1 otherwise
package pack

// offCode encodes the offset column (§4.7 C5).
func offCode(locs []Loc, j int) uint64 {
	l := locs[j]
	switch {
	case l.Enc == encGone:
		return 1
	case j > 0 && l.Enc != encGone && contiguous(locs[j-1], l):
		return 0
	default:
		return l.Off + 1
	}
}
