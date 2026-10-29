//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what coverHistory — dictionary content from samples: the best segment of each epoch (COVER-like), weakest first
package pack

import (
	"bytes"
	"sort"
)

// coverHistory picks up to size bytes. The samples are joined and cut into
// size/coverSegment epochs; each gives the window whose d-mers occur in the
// most samples (d-mers already picked count 0). Segments are ordered by score,
// best last, where zstd offsets are cheapest. Deterministic: maps are only
// looked up, never ranged over.
func coverHistory(samples [][]byte, size int) []byte {
	all := bytes.Join(samples, nil)
	epochs := min(size, len(all)) / coverSegment
	freq := dmerFreq(samples)
	var segs []segment
	for e := 0; e < epochs; e++ {
		lo, hi := e*len(all)/epochs, (e+1)*len(all)/epochs
		at, score := bestSegment(all, freq, lo, hi)
		if score == 0 {
			continue
		}
		b := all[at : at+coverSegment]
		segs = append(segs, segment{score: score, b: b})
		zeroDmers(freq, b)
	}
	sort.SliceStable(segs, func(i, j int) bool { return segs[i].score < segs[j].score })
	var out []byte
	for _, s := range segs {
		out = append(out, s.b...)
	}
	return out
}
