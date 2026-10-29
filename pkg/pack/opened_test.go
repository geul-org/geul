//ff:func feature=pack type=helper control=sequence
//ff:what opened — the latest generation of a local pack set, failing the test on error
package pack

import (
	"context"
	"testing"
)

// opened opens out through latest.txt.
func opened(t *testing.T, out string) *Reader {
	t.Helper()
	r, err := Open(context.Background(), FileStore{Root: out}, Options{})
	if err != nil {
		t.Fatal(err)
	}
	return r
}
