//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.At — one version of a document as a snapshot (0 = newest live); a stored diff is rebuilt and checked
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// At returns version v of document doc of issuer as a snapshot (snapshotOf:
// a stored diff is applied to its base and the content hash checked). v 0
// is the newest version that is neither deprecated nor deleted. A deleted
// version is not served (edge/PUBLICATION); an unknown one is ErrNotFound.
func (d Dir) At(issuer, doc []core.Word, v uint32) ([]byte, error) {
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return nil, err
	}
	e, ok := pickVersion(lg.Entries, v)
	if !ok {
		return nil, fmt.Errorf("%w: %s/%s@%s", ErrNotFound, k.Issuer, k.Doc, FormatVersion(v))
	}
	return d.snapshotOf(context.Background(), k, lg.Entries, e, 0)
}
