//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what addDmers — count each distinct 8-byte substring of one sample once
package pack

import "encoding/binary"

// addDmers adds 1 to freq for every distinct d-mer of s; seen is scratch.
func addDmers(freq map[uint64]uint32, seen map[uint64]bool, s []byte) {
	clear(seen)
	for i := 0; i+coverDmer <= len(s); i++ {
		v := binary.LittleEndian.Uint64(s[i:])
		if !seen[v] {
			seen[v] = true
			freq[v]++
		}
	}
}
