//ff:func feature=pack type=writer control=sequence
//ff:what Compact — a new generation holding only live records in fresh objects, one P and one H layer (§6 압축)
package pack

import (
	"context"
	"errors"
)

// Compact rewrites the current generation of out into new objects: the new
// OBJECTS lists only them, so older objects are no longer referenced (delete
// them after the retention period). The generation chain continues.
func Compact(ctx context.Context, out string, o BuildOptions) (BuildResult, error) {
	r, err := openPrev(ctx, out)
	if err == nil && r == nil {
		err = errors.New("geul/pack: nothing to compact (no latest.txt)")
	}
	if err != nil {
		return BuildResult{}, err
	}
	paths, err := r.Paths(ctx)
	if err != nil {
		return BuildResult{}, err
	}
	get := func(p string) ([]byte, error) { return r.Fetch(ctx, p) }
	o.Fresh = true
	return Write(ctx, out, Batch{Put: paths, Full: true, Get: get}, o)
}
