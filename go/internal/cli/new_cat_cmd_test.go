//ff:func feature=cli type=command control=sequence
//ff:what TestNewCatCmd — cat writes a version's bytes; a missing version fails
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestNewCatCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	out := h.must("cat", addr+"@2")
	ws, _ := core.ReadWords([]byte(out))
	if _, err := core.SplitDocument(ws); err != nil {
		t.Fatal(err)
	}
	if _, err := h.run("", "cat", addr+"@9"); err == nil {
		t.Fatal("want error")
	}
}
