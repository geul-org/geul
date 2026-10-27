//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.Current — a document's newest published version as a snapshot, with its version list's NEXT-LOCAL-ID; nil when it was never published
package resolve

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// Current returns the newest live version of document doc of issuer as a
// snapshot (a stored diff is applied to its base, snapshotOf) and the
// NEXT-LOCAL-ID of its version list (meta/lineage/NEXT-LOCAL-ID), for a
// producer that carries LOCAL IDs into the next version (delta.LocalIDs). A
// document never published gives nil, 0 and no error.
func (d Dir) Current(issuer, doc []core.Word) ([]byte, uint64, error) {
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return nil, 0, err
	}
	v, ok := newestLive(lg.Entries)
	switch {
	case !lg.Present:
		return nil, 0, nil
	case !ok:
		return nil, lg.Next, errors.New("geul/resolve: no live version")
	}
	b, err := d.snapshotOf(context.Background(), k, lg.Entries, v, 0)
	return b, lg.Next, err
}
