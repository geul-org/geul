//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterPlan — duplicates collapse, unchanged paths drop, changed ones stay; Get errors and sharing errors fail
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestWriterPlan(t *testing.T) {
	_, out := built(t, small())
	get := memGet(map[string][]byte{"issuers.geul": fixtureFile(t, "issuers.geul"), "x": {1}})
	w := testWriter(t, out, Batch{Put: []string{"x", "issuers.geul", "x"}, Get: get})
	if err := w.plan(); err != nil || len(w.puts) != 1 || w.puts[0].path != "x" || len(w.sums) != 2 {
		t.Fatal(err, w.puts)
	}
	w = testWriter(t, out, Batch{Put: []string{"y"}, Get: get})
	if err := w.plan(); err == nil {
		t.Fatal("get")
	}
	h := "h/0x" + strings.Repeat("0", 64) + ".geul"
	w = testWriter(t, out, Batch{Put: []string{h}, Get: memGet(map[string][]byte{h: {1, 2}})})
	if err := w.plan(); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
