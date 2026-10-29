//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestPackCompact — every set gets a fresh generation; a directory archive and missing objects fail
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackCompact(t *testing.T) {
	h := newHarness(t)
	h.doc()
	src := h.data
	h.packed()
	var out bytes.Buffer
	if err := packCompact(h.data, packOptions(), &out); err != nil || strings.Count(out.String(), "generation 2") != 2 {
		t.Fatal(err, out.String())
	}
	if err := packCompact(src, packOptions(), &out); err == nil {
		t.Fatal("directory archive")
	}
	objs, _ := filepath.Glob(filepath.Join(h.data, ".well-known", "geul", "*", "o", "*.gpk"))
	for _, o := range objs {
		os.Remove(o)
	}
	if err := packCompact(h.data, packOptions(), &out); err == nil {
		t.Fatal("missing objects")
	}
}
