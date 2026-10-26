//ff:func feature=cli type=helper control=sequence
//ff:what TestAppLoad — the persistent flags win over the environment
package cli

import "testing"

func TestAppLoad(t *testing.T) {
	h := newHarness(t)
	h.env["GEULFORMAT"] = "json"
	if out := h.must("env", "GEULFORMAT", "--format", "text"); out != "text\n" {
		t.Fatal(out)
	}
	if out := h.must("env", "GEULDATA", "--data", "/x"); out != "/x\n" {
		t.Fatal(out)
	}
}
