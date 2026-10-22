//ff:type feature=toulmin type=model
//ff:what Gate — a confidence gate (operator + integer percent)
package toulmin

// Gate is a confidence gate such as "if at least 75% certain". On the wire it is
// an operator word (bits 15-14, the rest 0) followed by an UNSIGNED-DECIMAL Value
// with the percent (GEUL t002 TOULMIN LINK).
type Gate struct {
	Op      GateOp
	Percent uint8
}

// gateMaxPercent is the highest legal confidence percent in a Gate.
const gateMaxPercent = 100
