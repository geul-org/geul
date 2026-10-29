//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveBatches — puts grouped by target set (h/ with the commit's Issuer), removals to every set that may have them; routing errors fail
package pack

import (
	"context"
	"testing"
)

func TestArchiveBatches(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	a.WriteFile("0x0001/0x0000000000000001/versions.geul", []byte("v"))
	a.WriteFile("h/new.geul", []byte("h"))
	a.Remove("issuers.geul")
	bs, err := a.batches(ctx)
	if err != nil || len(bs) != 2 {
		t.Fatal(bs, err)
	}
	b1 := bs[a.find("0x0001")]
	if len(b1.Put) != 2 || len(b1.Delete) != 0 {
		t.Fatal(b1)
	}
	if b0 := bs[a.find("0x0000")]; len(b0.Delete) != 1 || len(b0.Put) != 0 {
		t.Fatal(b0)
	}
	if got, err := b1.Get("h/new.geul"); err != nil || string(got) != "h" {
		t.Fatal(err)
	}
	a.find("0x0000").r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	a.WriteFile("0x0000/0x0000000000000008/versions.geul", nil)
	if _, err := a.batches(ctx); err == nil {
		t.Fatal("only")
	}
	a.find("0x0000").r.m.hLayers = []layer{{blk: block{kind: blkHInner, keys: []string{string(make([]byte, 32))}, locs: []Loc{{Obj: 99}}}}}
	a.puts = map[string][]byte{hPath(string(make([]byte, 32))): nil}
	if _, err := a.batches(ctx); err == nil {
		t.Fatal("target")
	}
}
