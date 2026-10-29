//ff:func feature=pack type=helper control=sequence
//ff:what built — a fixture archive and its first pack generation built with o
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

// built returns (archive root, pack out dir).
func built(t *testing.T, o BuildOptions) (string, string) {
	t.Helper()
	root := fixture(t)
	out := filepath.Join(t.TempDir(), "pack")
	if _, err := Build(context.Background(), root, out, o); err != nil {
		t.Fatal(err)
	}
	return root, out
}
