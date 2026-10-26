//ff:func feature=cli type=helper control=sequence
//ff:what TestAppJSON — text format writes nothing; json writes indented JSON
package cli

import (
	"strings"
	"testing"
)

func TestAppJSON(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if out := h.must("ls"); strings.HasPrefix(out, "[") {
		t.Fatal(out)
	}
	if out := h.must("ls", "--format", "json"); !strings.HasPrefix(out, "[\n  {") {
		t.Fatal(out)
	}
}
