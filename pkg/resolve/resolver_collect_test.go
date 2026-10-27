//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverCollect — B's pin (A v1) and C's pin (A v2) stay separate Nodes; a missing import fails
package resolve

import (
	"context"
	"testing"
)

func TestResolverCollect(t *testing.T) {
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "tree", Source: sampleTree(t)}}}}
	sel, err := r.collect(context.Background(), []Key{sampleKey(3, 0), sampleKey(3, 1)})
	if err != nil || len(sel.nodes) != 4 || sel.of(sampleKey(1, 1)) == nil || sel.of(sampleKey(1, 2)) == nil || sel.of(sampleKey(2, 0)) == nil {
		t.Fatalf("got %+v %v", sel, err)
	}
	if sel.of(sampleKey(3, 0)) != sel.of(sampleKey(3, 1)) {
		t.Fatal("latest C and pinned C v1 should share one Node")
	}
	if _, err := r.collect(context.Background(), []Key{sampleKey(9, 0)}); err == nil {
		t.Fatal("missing: want error")
	}
}
