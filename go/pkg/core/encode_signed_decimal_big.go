//ff:func feature=core type=encoder control=sequence
//ff:what EncodeSignedDecimalBig — a SIGNED-DECIMAL Value of any coefficient size: C×10^E
package core

import "math/big"

// EncodeSignedDecimalBig builds the SIGNED-DECIMAL Value C×10^E for a
// coefficient of any size (EncodeDecimalValue stops at int64): the 2-word
// exponent, then c in minimal two's complement WORDs. C and E are kept as
// given, so the quantum ("1.20" = 120×10^-2) stays.
func EncodeSignedDecimalBig(e int32, c *big.Int) Record {
	payload := append([]Word{Word(uint32(e) >> 16), Word(uint32(e))}, bigIntWords(c)...)
	return Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: payload}
}
