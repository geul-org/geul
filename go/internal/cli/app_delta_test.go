//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppDelta — a second argument that is not @v, a bad address, a missing version, a reordered version and an unwritable output fail
package cli

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/resolve"
)

func TestAppDelta(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	f := filepath.Join(h.data, "delta.geul")
	k, _ := h.key(addr)
	d := resolve.Dir{Root: h.data}
	v2, _ := d.At(k.Issuer, k.Doc, 2)
	ws, _ := core.ReadWords(v2)
	doc, _ := core.SplitDocument(ws)
	swapped, _ := core.EncodeDocument(doc.Head, []core.Record{doc.Body[1], doc.Body[0], doc.Body[2]})
	d.Publish(swapped)
	cases := [][]string{{addr + "@1", addr + "@2"}, {"x@1", "@2"}, {addr + "@1", "@9"}, {addr + "@2", "@3"}}
	for _, c := range cases {
		if _, err := h.run("", "delta", c[0], c[1], "-o", f); err == nil {
			t.Errorf("%v: want error", c)
		}
	}
	if _, err := h.run("", "delta", addr+"@1", "@2", "-o", h.data); err == nil {
		t.Fatal("output is a folder: want error")
	}
}
