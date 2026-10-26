//ff:func feature=cli type=command control=sequence
//ff:what TestNewEnvCmd — env lists, looks up, writes and removes settings; -w with -u fails
package cli

import (
	"strings"
	"testing"
)

func TestNewEnvCmd(t *testing.T) {
	h := newHarness(t)
	if out := h.must("env"); !strings.Contains(out, "GEULDATA="+h.data+"  # found") {
		t.Fatalf("list: %s", out)
	}
	h.must("env", "-w", "GEULISSUER=me")
	if out := h.must("env", "GEULISSUER"); out != "me\n" {
		t.Fatalf("written: %q", out)
	}
	h.must("env", "-u", "GEULISSUER")
	if out := h.must("env", "GEULISSUER"); out != "\n" {
		t.Fatalf("removed: %q", out)
	}
	if _, err := h.run("", "env", "-w", "-u", "GEULISSUER"); err == nil {
		t.Fatal("-w -u: want error")
	}
}
