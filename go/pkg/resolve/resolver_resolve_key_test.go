//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverResolveKey — pinned ignores AsOf, latest follows AsOf or the newest version
package resolve

import (
	"context"
	"testing"
)

func TestResolverResolveKey(t *testing.T) {
	ctx := context.Background()
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "law", Source: sampleLawTree(t)}}}}
	if n, err := r.resolveKey(ctx, sampleKey(5, 0)); err != nil || n.Version != 3 {
		t.Fatalf("newest: %+v %v", n, err)
	}
	r.AsOf = day(250)
	if n, err := r.resolveKey(ctx, sampleKey(5, 0)); err != nil || n.Version != 2 {
		t.Fatalf("as of: %+v %v", n, err)
	}
	if n, err := r.resolveKey(ctx, sampleKey(5, 1)); err != nil || n.Version != 1 {
		t.Fatalf("pinned: %+v %v", n, err)
	}
}
