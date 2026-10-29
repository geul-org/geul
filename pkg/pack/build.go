//ff:func feature=pack type=writer control=sequence
//ff:what Build — make the next generation of out from an archive directory: every file put, every vanished path deleted
package pack

import (
	"context"
	"os"
	"path/filepath"
)

// Build writes the archive directory root (`<root>/.well-known/geul/…`, the
// resolve.Dir layout) as the next generation of the pack set in out. Unchanged
// paths keep their records; the first build writes everything.
func Build(ctx context.Context, root, out string, o BuildOptions) (BuildResult, error) {
	paths, err := listTree(root)
	if err != nil {
		return BuildResult{}, err
	}
	base := filepath.Join(root, ".well-known", "geul")
	get := func(p string) ([]byte, error) { return os.ReadFile(filepath.Join(base, filepath.FromSlash(p))) }
	return Write(ctx, out, Batch{Put: paths, Full: true, Get: get}, o)
}
