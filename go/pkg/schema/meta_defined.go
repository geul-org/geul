//ff:func feature=schema type=rule control=selection
//ff:what metaDefined — whether a Meta codec is assigned in OVERVIEW
package schema

// metaDefined reports whether c is an assigned Meta codec (OVERVIEW Meta tables).
func metaDefined(c uint8) bool {
	switch {
	case c <= 0x02, c == 0x0E, c == 0x0F, c >= 0x10 && c <= 0x14, c == 0x16:
		return true
	case c >= 0x20 && c <= 0x28, c >= 0x30 && c <= 0x3C, c >= 0x40 && c <= 0x44, c >= 0x50 && c <= 0x55, c == 0x60:
		return true
	default:
		return false
	}
}
