//ff:func feature=schema type=rule control=selection
//ff:what numberProblem — why a numeric Value inside QUANTITY or FRACTION is not canonical: non-minimal integer or coefficient, short decimal, non-IEEE or non-finite BINARY-FLOAT, invalid FRACTION; "" when it is
package schema

import "github.com/geul-org/gogeul/pkg/core"

// numberProblem checks one numeric Value record held inside a QUANTITY or a
// FRACTION (value/numeric/QUANTITY, FRACTION): integers and coefficients in
// the fewest words (V3), a decimal with its 2-word exponent, a BINARY-FLOAT of
// 1/2/4/8 WORDs with a finite value, a FRACTION by fractionProblem. Other
// codecs are the decoders' form check and give "".
func numberProblem(r core.Record) string {
	switch r.Codec {
	case core.CodecSignedInteger, core.CodecUnsignedInteger:
		if r.Codec == core.CodecSignedInteger && !minimalSigned(r.Payload) || r.Codec == core.CodecUnsignedInteger && !minimalUnsigned(r.Payload) {
			return "최소 WORD가 아닌 정수"
		}
	case core.CodecSignedDecimal, core.CodecUnsignedDecimal:
		if len(r.Payload) < 3 || r.Codec == core.CodecSignedDecimal && !minimalSigned(r.Payload[2:]) || r.Codec == core.CodecUnsignedDecimal && !minimalUnsigned(r.Payload[2:]) {
			return "exponent 2 WORD + 최소 WORD coefficient가 아닌 10진수"
		}
	case core.CodecBinaryFloat:
		if !finiteFloat(r.Payload) {
			return "1/2/4/8 WORD 유한 BINARY-FLOAT이 아님(NaN·±Infinity 또는 폭)"
		}
	case core.CodecFraction:
		f, err := core.DecodeFraction(r)
		if err != nil {
			return err.Error()
		}
		return fractionProblem(f)
	}
	return ""
}
