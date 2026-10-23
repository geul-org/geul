//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverCollect — C pulls B and A; A's highest pin (v2) wins over B's v1; a missing import fails
package resolve

import (
	"context"
	"testing"
)

func TestResolverCollect(t *testing.T) {
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "tree", Source: sampleTree(t)}}}}
	sel, err := r.collect(context.Background(), []Key{sampleKey(3, 0), sampleKey(3, 1)})
	if err != nil || len(sel) != 3 || sel[sampleKey(1, 0).DocKey()].Version != 2 || sel[sampleKey(2, 0).DocKey()].Version != 1 {
		t.Fatalf("got %+v %v", sel, err)
	}
	if _, err := r.collect(context.Background(), []Key{sampleKey(9, 0)}); err == nil {
		t.Fatal("missing: want error")
	}
}
