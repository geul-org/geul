//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.readList — a document's version list document from the local tree; absent when never published
package resolve

import (
	"context"
	"errors"
)

// readList reads the version list of k; a document never published has none
// (Present false).
func (d Dir) readList(k DocKey) (versionLog, error) {
	b, err := d.Fetch(context.Background(), listPath(k))
	if errors.Is(err, ErrNotFound) {
		return versionLog{}, nil
	}
	if err != nil {
		return versionLog{}, err
	}
	return decodeVersionList(b, k)
}
