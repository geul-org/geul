//ff:func feature=pack type=decoder control=sequence
//ff:what parseHashHex — 64 lowercase hex digits (a manifest or object name) as 32 bytes
package pack

// parseHashHex decodes an object name.
func parseHashHex(s string) ([32]byte, bool) {
	var h [32]byte
	b, ok := parseHex("0x"+s, 64)
	copy(h[:], b)
	return h, ok
}
