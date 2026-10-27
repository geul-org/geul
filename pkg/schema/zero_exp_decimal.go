//ff:func feature=schema type=rule control=sequence
//ff:what zeroExpDecimal — whether a record is a SIGNED-/UNSIGNED-DECIMAL whose exponent E is 0
package schema

import "github.com/geul-org/gogeul/pkg/core"

// zeroExpDecimal reports whether r is a decimal Value with exponent 0 — a
// FRACTION component written that way is invalid (an integer original is an
// INTEGER record).
func zeroExpDecimal(r core.Record) bool {
	decimal := r.Codec == core.CodecSignedDecimal || r.Codec == core.CodecUnsignedDecimal
	return decimal && len(r.Payload) >= 2 && r.Payload[0] == 0 && r.Payload[1] == 0
}
