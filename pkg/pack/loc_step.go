//ff:func feature=pack type=decoder control=selection
//ff:what locStep — entry j's Loc from its C4–C7 values, checking deletion markers and contiguous offsets
package pack

import "math"

// locStep decodes one location. A deletion marker (leaf only) must be C4 = 0,
// C5 = 1, C6 = 0; C5 = 0 continues the previous entry in the same object and
// is invalid first or after a deletion marker.
func locStep(locs []Loc, j int, obj, off, n uint64, enc byte, leaf bool) (Loc, bool) {
	switch {
	case obj > math.MaxUint32:
		return Loc{}, false
	case enc == encGone:
		return Loc{Enc: encGone}, leaf && obj == 0 && off == 1 && n == 0
	case enc > encZstd:
		return Loc{}, false
	case off != 0:
		return Loc{Obj: uint32(obj), Off: off - 1, Len: n, Enc: enc}, true
	case j == 0 || locs[j-1].Enc == encGone || locs[j-1].Obj != uint32(obj):
		return Loc{}, false
	default:
		p := locs[j-1]
		return Loc{Obj: uint32(obj), Off: p.Off + p.Len, Len: n, Enc: enc}, true
	}
}
