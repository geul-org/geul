//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestFence — three backticks by default, longer than any inner run
package geulmd

import "testing"

func TestFence(t *testing.T) {
	if fence("a`b") != "```" || fence("x ```` y") != "`````" {
		t.Fatal("fence length mismatch")
	}
}
