//ff:func feature=pack type=rule control=sequence
//ff:what TestReaderVerifyDicts — a set's dictionaries load; a DICTS entry pointing at other bytes fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderVerifyDicts(t *testing.T) {
	out, _, _ := zstdSet(t, 100, &Zstd{DictSize: 4096})
	r := opened(t, out)
	if err := r.verifyDicts(context.Background()); err != nil || len(r.decs) != 1 {
		t.Fatal(err)
	}
	r = opened(t, out)
	r.m.dicts[0].loc.Off++
	if err := r.verifyDicts(context.Background()); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
