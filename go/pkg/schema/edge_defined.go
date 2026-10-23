//ff:func feature=schema type=rule control=selection
//ff:what edgeDefined — whether a direct Edge codec is assigned in OVERVIEW
package schema

// edgeDefined reports whether c is an assigned Edge codec (OVERVIEW Edge
// tables) or an ISSUER code.
func edgeDefined(c uint8) bool {
	switch {
	case c <= 0x02, c == 0x04, c == 0x05, c >= 0x08 && c <= 0x0B, c >= 0x10 && c <= 0x12, c >= 0x18 && c <= 0x1A:
		return true
	case c >= 0x20 && c <= 0x26, c >= 0x40 && c <= 0x59:
		return true
	case c >= 0x60 && c <= 0x65, c >= 0x68 && c <= 0x6A, c >= 0x70 && c <= 0x76, c == 0x78, c == 0x79, c >= 0x7C:
		return true
	default:
		return false
	}
}
