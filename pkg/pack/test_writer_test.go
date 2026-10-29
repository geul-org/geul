//ff:func feature=pack type=helper control=sequence
//ff:what testWriter — an opened writer over out with a never-failing file system and small options
package pack

import (
	"context"
	"testing"
)

// testWriter prepares a writer for out (which may be empty) and batch b.
func testWriter(t *testing.T, out string, b Batch) *writer {
	t.Helper()
	w := &writer{ctx: context.Background(), out: out, o: small().withDefaults(), fs: never(), b: b}
	if err := w.open(); err != nil {
		t.Fatal(err)
	}
	w.o.Codec, _ = w.followCodec()
	return w
}
