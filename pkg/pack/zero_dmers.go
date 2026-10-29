//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what zeroDmers — stop counting the d-mers of a chosen segment, so later picks bring new content
package pack

import "encoding/binary"

// zeroDmers sets the frequency of every d-mer of b to 0.
func zeroDmers(freq map[uint64]uint32, b []byte) {
	for i := 0; i+coverDmer <= len(b); i++ {
		freq[binary.LittleEndian.Uint64(b[i:])] = 0
	}
}
