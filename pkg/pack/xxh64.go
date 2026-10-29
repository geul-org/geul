//ff:func feature=pack type=helper control=sequence
//ff:what xxh64 — XXH64 with seed 0, the hash zstd's Content_Checksum keeps the low 32 bits of
package pack

// xxh64 computes XXH64(b, seed 0).
func xxh64(b []byte) uint64 {
	h, rest := xxP5, b
	if len(b) >= 32 {
		h, rest = xxStripes(b)
	}
	return xxTail(h+uint64(len(b)), rest)
}
