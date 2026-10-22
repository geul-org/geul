//ff:type feature=inspect type=model
//ff:what Value — a decoded Value record of a codec this library knows
package inspect

// Value is a decoded Value record. Codec is the codec as written (the odd
// member of a byte-string pair kept). Text holds UTF8-TEXT and BCP 47 tags,
// Uint an UNSIGNED-INTEGER, and Exp/Coef a SIGNED- or UNSIGNED-DECIMAL as
// written (Coef×10^Exp, quantum kept).
type Value struct {
	Codec uint8
	Text  string
	Uint  uint64
	Exp   int32
	Coef  int64
}
