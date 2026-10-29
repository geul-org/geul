//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Unpack — write every path of a generation back as files under `<root>/.well-known/geul/` (refuses an existing tree)
package pack

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// Unpack restores the resolve.Dir layout from r into root, which must not
// already hold `.well-known/geul/`. It returns the number of files written.
func Unpack(ctx context.Context, r *Reader, root string) (int, error) {
	base := filepath.Join(root, ".well-known", "geul")
	if _, err := os.Stat(base); err == nil {
		return 0, fmt.Errorf("geul/pack: %s already exists", base)
	}
	paths, err := r.Paths(ctx)
	if err != nil {
		return 0, err
	}
	for i, p := range paths {
		if err := unpackOne(ctx, r, base, p); err != nil {
			return i, err
		}
	}
	return len(paths), nil
}
