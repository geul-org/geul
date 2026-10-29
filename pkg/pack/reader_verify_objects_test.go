//ff:func feature=pack type=rule control=sequence
//ff:what TestReaderVerifyObjects — a built set passes; one flipped byte in a data pack fails (V5)
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderVerifyObjects(t *testing.T) {
	_, out := built(t, small())
	r := opened(t, out)
	if err := r.VerifyObjects(context.Background()); err != nil {
		t.Fatal(err)
	}
	flip(t, objPath(t, r, out, kindData), 100)
	if err := r.VerifyObjects(context.Background()); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
