//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverResolve — a root importing C gets one Node per document with A at v2 shared by B and C; bad root and missing import fail
package resolve

import (
	"context"
	"testing"
)

func TestResolverResolve(t *testing.T) {
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "tree", Source: sampleTree(t)}}}}
	g, err := r.Resolve(context.Background(), sampleDoc(7, "root", [3]uint32{1, 3, 0}))
	if err != nil || len(g.Docs) != 3 {
		t.Fatalf("got %+v %v", g, err)
	}
	c := g.Root.Imports[1]
	a := g.Docs[sampleKey(1, 0).DocKey()]
	if c == nil || c.Imports[2] != a || c.Imports[1].Imports[1] != a || a.Version != 2 || a.Imports == nil {
		t.Fatalf("links: %+v", g)
	}
	if _, err := r.Resolve(context.Background(), []byte{1}); err == nil {
		t.Fatal("bad root: want error")
	}
	if _, err := r.Resolve(context.Background(), sampleDoc(7, "root", [3]uint32{1, 9, 0})); err == nil {
		t.Fatal("missing import: want error")
	}
}
