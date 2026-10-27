//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.snapshotOf — a stored version as a snapshot: the file itself, or a diff applied to its base snapshot, checked against the list's content hash
package resolve

import (
	"context"
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/delta"
)

// maxDeltaChain bounds how many diffs a snapshot may stack (OVERVIEW 계보
// diff, DELTA-BASE 미결정 스냅숏 간격).
const maxDeltaChain = 64

// snapshotOf reads version e of k. A file with DELTA-BASE is a diff: its
// base version is rebuilt first (recursively, at most maxDeltaChain deep) and
// the diff applied (delta.Apply). The result's content hash must equal the
// list's hash for e.
func (d Dir) snapshotOf(ctx context.Context, k DocKey, es []VersionEntry, e VersionEntry, depth int) ([]byte, error) {
	b, err := d.Fetch(ctx, versionPath(k, e.Version))
	if err != nil {
		return nil, err
	}
	base, isDiff := deltaBaseOf(b)
	switch {
	case !isDiff:
		return checkHash(b, e.Hash)
	case depth >= maxDeltaChain:
		return nil, errors.New("geul/resolve: diff chain too long")
	}
	be, ok := entryOf(es, base)
	if !ok {
		return nil, fmt.Errorf("geul/resolve: diff base version %d not in the list", base)
	}
	bb, err := d.snapshotOf(ctx, k, es, be, depth+1)
	if err != nil {
		return nil, err
	}
	out, err := delta.Apply(bb, b)
	if err != nil {
		return nil, err
	}
	return checkHash(out, e.Hash)
}
