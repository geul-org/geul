//ff:func feature=resolve type=resolver control=sequence
//ff:what TestMemoryFetch — stored bytes, ErrNotFound when absent
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestMemoryFetch(t *testing.T) {
	m := Memory{"a": []byte("x")}
	if b, err := m.Fetch(context.Background(), "a"); err != nil || string(b) != "x" {
		t.Fatalf("hit: %q %v", b, err)
	}
	if _, err := m.Fetch(context.Background(), "b"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("miss: %v", err)
	}
}
