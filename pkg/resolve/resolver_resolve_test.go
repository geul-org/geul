//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverResolve — B keeps A v1 and C keeps A v2 in one graph; AsOf binds latest imports to the version in effect; bad root and missing import fail
package resolve

import (
	"context"
	"testing"
)

func TestResolverResolve(t *testing.T) {
	ctx := context.Background()
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "tree", Source: sampleTree(t)}}}}
	g, err := r.Resolve(ctx, sampleDoc(7, "root", [3]uint32{1, 3, 0}))
	if err != nil || len(g.Docs) != 4 {
		t.Fatalf("got %+v %v", g, err)
	}
	c := g.Root.Imports[1]
	if c == nil || c.Imports[2].Version != 2 || c.Imports[1].Imports[1].Version != 1 || c.Imports[2] == c.Imports[1].Imports[1] {
		t.Fatalf("bindings: %+v", g)
	}
	if _, err := r.Resolve(ctx, []byte{1}); err == nil {
		t.Fatal("bad root: want error")
	}
	if _, err := r.Resolve(ctx, sampleDoc(7, "root", [3]uint32{1, 9, 0})); err == nil {
		t.Fatal("missing import: want error")
	}
	law := &Resolver{Config: Config{Proxy: []Entry{{Name: "law", Source: sampleLawTree(t)}}}, AsOf: day(250)}
	g, err = law.Resolve(ctx, sampleDoc(7, "root", [3]uint32{1, 6, 0}))
	if err != nil || g.Root.Imports[1].Imports[1].Version != 2 {
		t.Fatalf("as of: %+v %v", g, err)
	}
}
