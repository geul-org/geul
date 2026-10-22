//ff:type feature=core type=model
//ff:what Value codec constants — the GEUL t002 direct Value codes TANGEUL embeds
package core

// Direct Value codecs used by the TANGEUL profile (GEUL t002 OVERVIEW Value).
// Variable-length byte strings come in even/odd pairs: the low bit is 1 when
// the byte count is odd and the last word carries a 0x00 pad byte.
const (
	// CodecUnsignedInteger is a minimal-word big-endian unsigned integer.
	CodecUnsignedInteger uint8 = 0x02
	// CodecSignedDecimal is a signed 32-bit exponent E plus a minimal-word
	// two's-complement coefficient C, value C×10^E.
	CodecSignedDecimal uint8 = 0x10
	// CodecUnsignedDecimal is a signed 32-bit exponent E plus a minimal-word
	// unsigned coefficient C, value C×10^E.
	CodecUnsignedDecimal uint8 = 0x11
	// CodecUTF8TextEven is valid UTF-8 without implied normalization (even bytes).
	CodecUTF8TextEven uint8 = 0x40
	// CodecLanguageBCP47Even is a canonical BCP 47 language tag (even bytes).
	CodecLanguageBCP47Even uint8 = 0x46
)
