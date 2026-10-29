//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what statSamples — the samples a dictionary's entropy tables are fitted on: the first dictStatCount, each cut to one zstd block
package pack

// statSamples returns up to dictStatCount samples, each at most maxBlock
// bytes. klauspost's BuildDict encodes every content sample as a single
// block and panics on one larger than a block (128 KiB): a 244 KB codec spec
// document (0x0000/0x0000000000000848 v2) or a big Wikidata item broke
// training. Cutting keeps the result deterministic.
func statSamples(samples [][]byte) [][]byte {
	out := make([][]byte, 0, min(len(samples), dictStatCount))
	for _, s := range samples[:min(len(samples), dictStatCount)] {
		out = append(out, s[:min(len(s), maxBlock)])
	}
	return out
}
