//ff:func feature=resolve type=resolver control=sequence
//ff:what TestRegistryFetch — always ErrNotFound for now
package resolve

import (
	"context"
	"errors"
	"testing"
)

func TestRegistryFetch(t *testing.T) {
	if _, err := (Registry{}).Fetch(context.Background(), "x"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("got %v", err)
	}
}
