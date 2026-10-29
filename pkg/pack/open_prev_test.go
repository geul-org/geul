//ff:func feature=pack type=writer control=sequence
//ff:what TestOpenPrev — nil without latest.txt, the generation with it
package pack

import (
	"context"
	"testing"
)

func TestOpenPrev(t *testing.T) {
	if r, err := openPrev(context.Background(), t.TempDir(), nil); r != nil || err != nil {
		t.Fatal(r, err)
	}
	_, out := built(t, small())
	if r, err := openPrev(context.Background(), out, nil); err != nil || r.Generation() != 1 {
		t.Fatal(err)
	}
}
