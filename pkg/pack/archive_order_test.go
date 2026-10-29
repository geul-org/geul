//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveOrder — Issuer paths search their own set then the top set; other paths every opened set; sets without a generation are skipped
package pack

import "testing"

func TestArchiveOrder(t *testing.T) {
	r := &Reader{}
	top, i0, i1, empty := &packSet{r: r}, &packSet{name: "0x0000", r: r}, &packSet{name: "0x0001", r: r}, &packSet{name: "0x0002"}
	a := &Archive{sets: []*packSet{top, i0, i1, empty}}
	if o := a.order("0x0001/0x0000000000000001/versions.geul"); len(o) != 2 || o[0] != i1 || o[1] != top {
		t.Fatal(o)
	}
	if o := a.order("h/x"); len(o) != 3 || o[0] != top {
		t.Fatal(o)
	}
}
