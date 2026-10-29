//ff:func feature=pack type=helper control=sequence
//ff:what zstdSet — a pack set of synth records written with codec c (nil follows), and its Get
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

// zstdSet writes synth(n) to a new set and returns (out, paths, get).
func zstdSet(t *testing.T, n int, c Codec) (string, []string, func(string) ([]byte, error)) {
	t.Helper()
	paths, get := synth(n)
	out := filepath.Join(t.TempDir(), "s")
	o := small()
	o.Codec = c
	if _, err := Write(context.Background(), out, Batch{Put: paths, Get: get}, o); err != nil {
		t.Fatal(err)
	}
	return out, paths, get
}
