//ff:func feature=toulmin type=model control=iteration dimension=1
//ff:what TestControlCode — each ControlCode constant equals the logical code
package toulmin

import "testing"

func TestControlCode(t *testing.T) {
	cases := []struct {
		got  ControlCode
		want uint16
	}{
		{ControlSection, 0x03},
		{ControlCaseBegin, 0x04},
		{ControlCaseEnd, 0x05},
		{ControlProse, 0x06},
	}
	for _, c := range cases {
		if uint16(c.got) != c.want {
			t.Errorf("ControlCode: got 0x%02X, want 0x%02X", uint16(c.got), c.want)
		}
	}
}
