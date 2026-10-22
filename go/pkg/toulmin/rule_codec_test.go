//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestRuleCodec — the three rule roles and the unknown-role error
package toulmin

import "testing"

func TestRuleCodec(t *testing.T) {
	g, _ := ruleCodec(RoleGeneral)
	c, _ := ruleCodec(RoleCounter)
	e, _ := ruleCodec(RoleExcept)
	if g.Codec != 0x63 || c.Codec != 0x64 || e.Codec != 0x65 {
		t.Fatalf("rule codecs: %+v %+v %+v", g, c, e)
	}
	if _, err := ruleCodec(3); err == nil {
		t.Fatal("role 3 must fail")
	}
}
