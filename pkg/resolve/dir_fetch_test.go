//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirFetch — a written file reads back, a missing one is ErrNotFound, a directory is another error
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestDirFetch(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	if err := d.write("a/b.geul", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if b, err := d.Fetch(context.Background(), "a/b.geul"); err != nil || string(b) != "x" {
		t.Fatalf("hit: %q %v", b, err)
	}
	if _, err := d.Fetch(context.Background(), "a/c.geul"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("miss: %v", err)
	}
	if _, err := d.Fetch(context.Background(), "a"); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatalf("dir: %v", err)
	}
}
