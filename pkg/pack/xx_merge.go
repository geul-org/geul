//ff:func feature=pack type=helper control=sequence
//ff:what xxMerge — fold one XXH64 accumulator into the hash
package pack

// xxMerge merges accumulator v into h.
func xxMerge(h, v uint64) uint64 {
	h ^= xxRound(0, v)
	return h*xxP1 + xxP4
}
