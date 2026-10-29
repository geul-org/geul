//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderHas — P, H and named paths without reading records; unknown Issuers and absent paths are false; index errors are returned
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderHas(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	for _, p := range []string{"0x0000/0x0000000000000845/versions.geul", "h/0x009011a66aa763a58d70f493caf343936c1c909624dee0ce8950b0727199c2af.geul", "issuers.geul"} {
		if ok, err := r.Has(ctx, p); !ok || err != nil {
			t.Fatal(p, err)
		}
	}
	for _, p := range []string{"0x0001/0x0000000000000845/versions.geul", "0x0000/0x0000000000000846/versions.geul", "zz"} {
		if ok, err := r.Has(ctx, p); ok || err != nil {
			t.Fatal(p, err)
		}
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Has(ctx, "0x0000/0x0000000000000845/versions.geul"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
