//ff:func feature=cli type=command control=sequence
//ff:what TestNewResolveCmd — an IMPORT bound to a version found in the data folder
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewResolveCmd(t *testing.T) {
	h := newHarness(t)
	h.env["GEULPROXY"] = "off"
	addr, _ := h.doc()
	f := filepath.Join(h.data, "b.geul")
	h.must("new", "me/0x0000000000000002", "-o", f)
	h.must("meta", f, "IMPORT-DOC", addr+"@1", "0x0001")
	out := h.must("resolve", f)
	if !strings.HasSuffix(out, "/0x0000000000000001@1 file://"+h.data+"\n") {
		t.Fatalf("%q", out)
	}
}
