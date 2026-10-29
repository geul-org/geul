//ff:func feature=pack type=decoder control=sequence
//ff:what windowSize — the window a Window_Descriptor byte gives
package pack

// windowSize decodes Exponent (high 5 bits) and Mantissa (low 3 bits).
func windowSize(wd byte) uint64 {
	base := uint64(1) << (10 + uint(wd>>3))
	return base + base/8*uint64(wd&7)
}
