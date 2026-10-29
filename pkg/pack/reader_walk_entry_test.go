//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderWalkEntry — leaf entries call fn; internal entries descend (and report bad children)
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderWalkEntry(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	var got string
	leaf := block{kind: blkHLeaf, keys: []string{"a"}, locs: []Loc{{}}}
	if err := r.walkEntry(ctx, leaf, 0, "", "", 0, func(k string, _ Loc) error { got = k; return nil }); err != nil || got != "a" {
		t.Fatal(err)
	}
	in := block{kind: blkHInner, keys: []string{"a"}, locs: []Loc{{Obj: 9}}}
	if err := r.walkEntry(ctx, in, 0, "", "", 0, nil); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
