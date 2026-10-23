//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.readList — a document's version list from the local tree, empty when there is none
package resolve

import (
	"context"
	"errors"
)

// readList reads d's version list for k; a document never published has none.
func (d Dir) readList(k DocKey) ([]VersionEntry, error) {
	b, err := d.Fetch(context.Background(), listPath(k))
	if errors.Is(err, ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return parseVersionList(b)
}
