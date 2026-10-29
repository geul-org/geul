//ff:func feature=pack type=encoder control=selection
//ff:what pKeyDelta — the C1·C2·C3 column values of a P key given the previous key (§4.7)
package pack

// pKeyDelta encodes key j against the previous key: the first key is raw; a
// new Issuer stores its difference with D and S raw; a new document stores 0,
// the D difference and S raw; otherwise 0, 0 and the S difference.
func pKeyDelta(j int, pi uint32, pd uint64, ps uint32, i uint32, d uint64, s uint32) (uint64, uint64, uint64) {
	switch {
	case j == 0:
		return uint64(i), d, uint64(s)
	case i != pi:
		return uint64(i - pi), d, uint64(s)
	case d != pd:
		return 0, d - pd, uint64(s)
	default:
		return 0, 0, uint64(s - ps)
	}
}
