//ff:func feature=pack type=decoder control=selection
//ff:what pKeyStep — the next P key from the previous one and its C1·C2·C3 values, false on overflow or no increase
package pack

import "math"

// pKeyStep inverts pKeyDelta. It fails when a field overflows its width or
// the key does not increase (C1 = C2 = 0 needs C3 ≥ 1).
func pKeyStep(j int, pi uint32, pd uint64, ps uint32, a, b, c uint64) (uint32, uint64, uint32, bool) {
	switch {
	case j == 0:
		return uint32(a), b, uint32(c), a <= math.MaxUint32 && c <= math.MaxUint32
	case a != 0:
		return pi + uint32(a), b, uint32(c), a <= math.MaxUint32-uint64(pi) && c <= math.MaxUint32
	case b != 0:
		return pi, pd + b, uint32(c), b <= math.MaxUint64-pd && c <= math.MaxUint32
	default:
		return pi, pd, ps + uint32(c), c >= 1 && c <= math.MaxUint32-uint64(ps)
	}
}
