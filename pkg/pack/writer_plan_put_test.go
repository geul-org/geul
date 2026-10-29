//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterPlanPut — invalid paths, Get errors, unreadable previous records, unchanged (nil) and changed puts
package pack

import (
	"os"
	"strings"
	"testing"
)

func TestWriterPlanPut(t *testing.T) {
	_, out := built(t, small())
	doc := fixtureFile(t, "issuers.geul")
	get := memGet(map[string][]byte{"issuers.geul": doc, "0x0000/documents.geul": {1, 2}})
	w := testWriter(t, out, Batch{Get: get})
	w.sums = map[string][32]byte{}
	for _, p := range []string{"", strings.Repeat("a", 0x10000), "none"} {
		if _, err := w.planPut(p); err == nil {
			t.Fatal(len(p))
		}
	}
	if pu, err := w.planPut("issuers.geul"); pu != nil || err != nil {
		t.Fatal(pu, err)
	}
	pu, err := w.planPut("0x0000/documents.geul")
	if err != nil || pu == nil || pu.chOK || pu.ref.space != spaceNamed {
		t.Fatal(pu, err)
	}
	os.Remove(objPath(t, w.prev, out, kindData))
	if _, err := w.planPut("issuers.geul"); err == nil {
		t.Fatal("unreadable previous record")
	}
}
