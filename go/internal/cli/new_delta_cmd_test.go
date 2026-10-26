//ff:func feature=cli type=command control=sequence
//ff:what TestNewDeltaCmd — the diff document rebuilds version 2; -o is required
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewDeltaCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	f := filepath.Join(h.data, "delta.geul")
	h.must("delta", addr+"@1", "@2", "-o", f)
	if out := h.must("show", f); !strings.Contains(out, "DELTA-BASE 0x00000001") || !strings.Contains(out, "#3 ALIAS") {
		t.Fatal(out)
	}
	if _, err := h.run("", "delta", addr+"@1", "@2"); err == nil {
		t.Fatal("no -o: want error")
	}
}
