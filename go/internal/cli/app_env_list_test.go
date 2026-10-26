//ff:func feature=cli type=command control=sequence
//ff:what TestAppEnvList — text with sources, unset keys without one; JSON with value and source
package cli

import (
	"strings"
	"testing"
)

func TestAppEnvList(t *testing.T) {
	h := newHarness(t)
	out := h.must("env")
	if !strings.Contains(out, "GEULFORMAT=text  # default\n") || !strings.Contains(out, "GEULISSUER=\n") {
		t.Fatal(out)
	}
	if out := h.must("env", "--format", "json"); !strings.Contains(out, `"Source": "found"`) {
		t.Fatal(out)
	}
}
