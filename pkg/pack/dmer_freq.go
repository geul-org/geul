//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what dmerFreq — in how many samples each 8-byte substring occurs
package pack

// dmerFreq is the document frequency of every d-mer.
func dmerFreq(samples [][]byte) map[uint64]uint32 {
	freq, seen := map[uint64]uint32{}, map[uint64]bool{}
	for _, s := range samples {
		addDmers(freq, seen, s)
	}
	return freq
}
