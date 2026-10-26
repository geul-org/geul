//ff:func feature=cli type=command control=sequence
//ff:what TestAppEnvRun — each mode: both flags fail, write, values, list
package cli

import (
	"strings"
	"testing"
)

func TestAppEnvRun(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "env", "-w", "-u", "X"); err == nil {
		t.Fatal("both: want error")
	}
	if _, err := h.run("", "env", "-w", "GEULFORMAT=json"); err != nil {
		t.Fatal(err)
	}
	if out := h.must("env", "GEULFORMAT"); out != "json\n" {
		t.Fatal(out)
	}
	if out := h.must("env"); !strings.HasPrefix(out, "{") {
		t.Fatal(out)
	}
}
