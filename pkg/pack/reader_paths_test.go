//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderPaths — exactly the fixture's files, sorted; index errors fail
package pack

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestReaderPaths(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	want, _ := listTree("testdata/archive")
	if got, err := r.Paths(ctx); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatal(got, err)
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Paths(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
