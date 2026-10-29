//ff:func feature=pack type=writer control=sequence
//ff:what streamTestNew — a new Stream in a temp folder with the test options but no dictionary samples (codec as given, nil for Zstd)
package pack

import (
	"path/filepath"
	"testing"
)

// streamTestNew starts a stream in a fresh folder with streamTestOptions
// minus the samples and with codec c (nil: Zstd).
func streamTestNew(t *testing.T, c Codec) *Stream {
	t.Helper()
	o := streamTestOptions()
	o.Samples, o.Build.Codec = nil, c
	s, err := NewStream(filepath.Join(t.TempDir(), "set"), o)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
