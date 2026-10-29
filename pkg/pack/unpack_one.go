//ff:func feature=pack type=writer control=sequence
//ff:what unpackOne — fetch one path and write it as a file under base
package pack

import (
	"context"
	"os"
	"path/filepath"
)

// unpackOne writes path's bytes to base/path.
func unpackOne(ctx context.Context, r *Reader, base, path string) error {
	b, err := r.Fetch(ctx, path)
	if err != nil {
		return err
	}
	dst := filepath.Join(base, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dst, b, 0o644)
}
