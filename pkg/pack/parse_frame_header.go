//ff:func feature=pack type=decoder control=sequence
//ff:what parseFrameHeader — magic, descriptor, window, dictionary ID and content size of a zstd frame
package pack

// parseFrameHeader reads a zstd frame header; a wrong magic or the reserved
// descriptor bit set is ErrCorrupt.
func parseFrameHeader(b []byte) (frameHeader, error) {
	c := &cursor{b: b}
	magic := c.u32()
	fhd := c.u8()
	if c.bad || magic != zstdMagic || fhd&0x08 != 0 {
		return frameHeader{}, corrupt("zstd frame header")
	}
	single := fhd&0x20 != 0
	h := frameHeader{checksum: fhd&0x04 != 0}
	if !single {
		h.window = windowSize(c.u8())
	}
	h.dictID = uint32(readLE(c.take(uint64([4]int{0, 1, 2, 4}[fhd&3]))))
	n := fcsSize(fhd>>6, single)
	h.hasFCS = n > 0
	h.fcs = readLE(c.take(uint64(n)))
	if n == 2 {
		h.fcs += 256
	}
	if single {
		h.window = h.fcs
	}
	if c.bad {
		return frameHeader{}, corrupt("zstd frame header")
	}
	h.size = len(b) - len(c.b)
	return h, nil
}
