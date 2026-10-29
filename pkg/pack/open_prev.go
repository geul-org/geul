//ff:func feature=pack type=writer control=sequence
//ff:what openPrev — the current generation of out, nil when out has no latest.txt yet
package pack

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// openPrev opens out's latest generation for a writer.
func openPrev(ctx context.Context, out string) (*Reader, error) {
	if _, err := os.Stat(filepath.Join(out, latestFile)); errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	return Open(ctx, FileStore{Root: out}, Options{})
}
