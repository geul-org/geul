package schema

// Value codecs accepted inside Meta (M4~M6).
var (
	isUTF8Text = func(c uint8) bool { return c == 0x40 || c == 0x41 }
	isLanguage = func(c uint8) bool { return c >= 0x46 && c <= 0x48 }
	isTime     = func(c uint8) bool { return c == 0x20 || c == 0x22 }
	isDigest   = func(c uint8) bool { return c == 0x61 }
)
