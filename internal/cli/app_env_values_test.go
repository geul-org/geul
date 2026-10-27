//ff:func feature=cli type=command control=sequence
//ff:what TestAppEnvValues — one value per key in order
package cli

import "testing"

func TestAppEnvValues(t *testing.T) {
	h := newHarness(t)
	if out := h.must("env", "GEULFORMAT", "GEULREMOTE"); out != "text\nhttps://geul.org\n" {
		t.Fatal(out)
	}
}
