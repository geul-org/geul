//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSourceList — a published list parses; a missing list is ErrNotFound; a bad list is ErrVerify
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestSourceList(t *testing.T) {
	ctx := context.Background()
	d := sampleKey(1, 0).DocKey()
	if l, err := sourceList(ctx, Entry{Source: sampleTree(t)}, d); err != nil || len(l) != 2 {
		t.Fatalf("list: %+v %v", l, err)
	}
	if _, err := sourceList(ctx, Entry{Source: Memory{}}, d); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
	if _, err := sourceList(ctx, Entry{Source: Memory{listPath(d): []byte("junk")}}, d); !errors.Is(err, ErrVerify) {
		t.Fatalf("bad: %v", err)
	}
}
