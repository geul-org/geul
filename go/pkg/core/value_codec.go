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
	// CodecGregorianDate is a fixed 2-word signed int32 day count from 1970-01-01.
	CodecGregorianDate uint8 = 0x20
	// CodecPosixSeconds is a fixed 4-word signed int64 second count from the Unix epoch.
	CodecPosixSeconds uint8 = 0x22
	// CodecEDTFEven is an EDTF (ISO 8601-2) date string, ASCII (even bytes;
	// 0x25 is the odd member).
	CodecEDTFEven uint8 = 0x24
	// CodecGEDCOMDateEven is a GEDCOM 7 date value string (calendar keyword,
	// e.g. JULIAN), ASCII (even bytes; 0x27 is the odd member).
	CodecGEDCOMDateEven uint8 = 0x26
	// CodecSHA256Digest is a fixed 16-word SHA-256 digest.
	CodecSHA256Digest uint8 = 0x61
	// CodecUTF8TextEven is valid UTF-8 without implied normalization (even bytes).
	CodecUTF8TextEven uint8 = 0x40
	// CodecLanguageBCP47Even is a canonical BCP 47 language tag (even bytes).
	CodecLanguageBCP47Even uint8 = 0x46
	// CodecUTF8LanguageTextEven is a language-tagged text: a language Value
	// record, then UTF-8 bytes (even text bytes; 0x4B is the odd member).
	CodecUTF8LanguageTextEven uint8 = 0x4A
)
