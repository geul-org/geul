//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppEnvWrite — several KEY=VALUE are written; -w needs KEY=VALUE; an unknown key fails
package cli

import "testing"

func TestAppEnvWrite(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "env", "-w", "GEULISSUER=me", "GEULREMOTE=https://x"); err != nil {
		t.Fatal(err)
	}
	if out := h.must("env", "GEULREMOTE"); out != "https://x\n" {
		t.Fatal(out)
	}
	for _, arg := range []string{"GEULDATA", "GEULDATA="} {
		if _, err := h.run("", "env", "-w", arg); err == nil {
			t.Errorf("%s: want error", arg)
		}
	}
	if _, err := h.run("", "env", "-w", "NOPE=1"); err == nil {
		t.Fatal("unknown key: want error")
	}
}
