//ff:func feature=resolve type=resolver control=sequence
//ff:what TestResolverFetch — pinned and newest versions through the source list
package resolve

import (
	"context"
	"testing"
)

func TestResolverFetch(t *testing.T) {
	r := &Resolver{Config: Config{Proxy: []Entry{{Name: "tree", Source: sampleTree(t)}}}}
	if n, err := r.Fetch(context.Background(), sampleKey(1, 1)); err != nil || n.Version != 1 {
		t.Fatalf("pinned: %+v %v", n, err)
	}
	if n, err := r.Fetch(context.Background(), sampleKey(1, 0)); err != nil || n.Version != 2 {
		t.Fatalf("newest: %+v %v", n, err)
	}
}
