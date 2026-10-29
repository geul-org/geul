//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveIndex — sorted union of set paths with staged puts and removals; built once; set errors fail
package pack

import (
	"testing"
)

func TestArchiveIndex(t *testing.T) {
	a, _ := testArchive(t)
	a.sets = append(a.sets, &packSet{name: "0x0009"})
	a.puts["a.geul"] = nil
	a.dels["issuers.geul"] = true
	idx, err := a.index()
	if err != nil || len(idx) != 18 || idx[len(idx)-1] != "h/0xe12a5f337013cbb5a3c5e4ba8381e114965f925a721124033a9c20dffde9d925.geul" || idx[0] != "0x0000/0x0000000000000008/0x00000001.geul" {
		t.Fatal(idx, err)
	}
	if again, _ := a.index(); &again[0] != &idx[0] {
		t.Fatal("rebuilt")
	}
	b, _ := testArchive(t)
	b.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := b.index(); err == nil {
		t.Fatal("broken")
	}
}
