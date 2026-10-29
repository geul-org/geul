//ff:func feature=pack type=helper control=sequence
//ff:what Zstd.dictSize — the size of a newly trained dictionary (0 is DefaultDictSize)
package pack

// dictSize is z.DictSize or DefaultDictSize; callers skip training when it is below 0.
func (z *Zstd) dictSize() int {
	return defaultInt(z.DictSize, 1, DefaultDictSize)
}
