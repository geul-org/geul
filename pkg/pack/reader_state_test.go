//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderState — inventories alone rebuild the same state as the index (V3); a broken data pack fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderState(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	st, err := r.State(ctx)
	if err != nil {
		t.Fatal(err)
	}
	idx, _ := r.entries(ctx)
	if sameState(st, idx) != nil {
		t.Fatal("V3")
	}
	flip(t, objPath(t, r, out, kindData), 0)
	if _, err := r.State(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
