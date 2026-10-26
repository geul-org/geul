//ff:func feature=cli type=command control=sequence
//ff:what TestNewLogCmd — log needs one document and lists its versions
package cli

import (
	"strings"
	"testing"
)

func TestNewLogCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("log", addr); !strings.HasPrefix(out, "1\t") || !strings.HasSuffix(out, "NEXT-LOCAL-ID 4\n") {
		t.Fatal(out)
	}
	if _, err := h.run("", "log"); err == nil {
		t.Fatal("no document: want error")
	}
}
