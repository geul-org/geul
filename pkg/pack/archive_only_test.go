//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveOnly — the one set Issuer paths go to; nil for none or several; routing errors pass up
package pack

import (
	"context"
	"testing"
)

func TestArchiveOnly(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	one := []string{"h/x", "0x0000/0x0000000000000001/versions.geul", "0x0000/documents.geul"}
	if s, err := a.only(ctx, one); err != nil || s != a.sets[0] {
		t.Fatal(s, err)
	}
	if s, _ := a.only(ctx, append(one, "0x0001/0x0000000000000001/versions.geul")); s != nil {
		t.Fatal("several")
	}
	if s, _ := a.only(ctx, []string{"h/x"}); s != nil {
		t.Fatal("none")
	}
	{
		n := "0x0000/0x0000000000000008/versions.geul"
		a.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
		if _, err := a.only(ctx, []string{n}); err == nil {
			t.Fatal("broken")
		}
	}
}
