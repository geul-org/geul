//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what bestSegment — the window of coverSegment bytes in [lo, hi) whose d-mers have the largest summed frequency
package pack

import "encoding/binary"

// bestSegment returns the window start and its score (0 when nothing scores).
func bestSegment(all []byte, freq map[uint64]uint32, lo, hi int) (int, uint64) {
	m := coverSegment - coverDmer + 1 // d-mers in one window
	var sum, best uint64
	at := lo
	for p := lo; p+coverDmer <= hi; p++ {
		sum += uint64(freq[binary.LittleEndian.Uint64(all[p:])])
		if p-lo >= m {
			sum -= uint64(freq[binary.LittleEndian.Uint64(all[p-m:])])
		}
		if p-lo >= m-1 && sum > best {
			best, at = sum, p-m+1
		}
	}
	return at, best
}
