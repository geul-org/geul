//ff:type feature=inspect type=model
//ff:what Value — a decoded Value record of a codec this library knows
package inspect

import "time"

// Value is a decoded Value record. Codec is the codec as written (the odd
// member of a byte-string pair kept). Text holds UTF8-TEXT and BCP 47 tags and
// a SHA-256-DIGEST as `0x` hex, Time a GREGORIAN-DATE or
// POSIX-TIMESTAMP-SECONDS (UTC), Uint an UNSIGNED-INTEGER, and Exp/Coef a
// SIGNED- or UNSIGNED-DECIMAL as written (Coef×10^Exp, quantum kept).
type Value struct {
	Codec uint8
	Text  string
	Time  time.Time
	Uint  uint64
	Exp   int32
	Coef  int64
}
