//ff:func feature=toulmin type=encoder control=selection
//ff:what ruleCodec — RULE/COUNTER/EXCEPT-DECLARE by rule role
package toulmin

import "fmt"

// ruleCodec maps a rule role to its declaration codec: general → RULE-DECLARE,
// counter → COUNTER-DECLARE, except → EXCEPT-DECLARE.
func ruleCodec(role RoleCode) (CodecRef, error) {
	switch role {
	case RoleGeneral:
		return CodecRef{Codec: CodecRuleDeclare}, nil
	case RoleCounter:
		return CodecRef{Codec: CodecCounterDeclare}, nil
	case RoleExcept:
		return CodecRef{Codec: CodecExceptDeclare}, nil
	default:
		return CodecRef{}, fmt.Errorf("geul/toulmin: unknown rule role %d", uint16(role))
	}
}
