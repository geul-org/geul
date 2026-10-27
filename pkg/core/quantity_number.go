//ff:func feature=core type=helper control=selection
//ff:what quantityNumber — whether a record is one of the numeric Values a QUANTITY amount or bound takes
package core

// quantityNumber reports whether x is a Value of the closed list a QUANTITY
// amount, lower and upper take (value/numeric/QUANTITY): SIGNED-INTEGER,
// UNSIGNED-INTEGER, BINARY-FLOAT, FRACTION, SIGNED-DECIMAL, UNSIGNED-DECIMAL.
func quantityNumber(x Record) bool {
	if x.Kind != KindValue {
		return false
	}
	switch x.Codec {
	case CodecSignedInteger, CodecUnsignedInteger, CodecBinaryFloat, CodecFraction, CodecSignedDecimal, CodecUnsignedDecimal:
		return true
	default:
		return false
	}
}
