//ff:func feature=pack type=encoder control=sequence
//ff:what appendRawBlock — a zstd raw block: 3-byte header (Last_Block, type 0, size) and the bytes
package pack

// appendRawBlock appends b (at most 128 KiB) as one raw block.
func appendRawBlock(out, b []byte, last bool) []byte {
	h := uint32(len(b)) << 3
	if last {
		h |= 1
	}
	out = append(out, byte(h), byte(h>>8), byte(h>>16))
	return append(out, b...)
}
