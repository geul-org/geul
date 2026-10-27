//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDecimalValue — a SIGNED-DECIMAL or UNSIGNED-DECIMAL Value record for C×10^E
package core

import "errors"

// EncodeDecimalValue builds the decimal Value record for c×10^e: the signed
// 32-bit exponent in two words, then the coefficient in minimal words —
// two's complement for SIGNED-DECIMAL (signed), unsigned magnitude for
// UNSIGNED-DECIMAL. A negative coefficient under UNSIGNED-DECIMAL is an error.
func EncodeDecimalValue(signed bool, e int32, c int64) (Record, error) {
	payload := []Word{Word(uint32(e) >> 16), Word(uint32(e))}
	if signed {
		return Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: append(payload, intWords(c)...)}, nil
	}
	if c < 0 {
		return Record{}, errors.New("geul/core: UNSIGNED-DECIMAL coefficient is negative")
	}
	return Record{Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: append(payload, uintWords(uint64(c))...)}, nil
}
