//ff:func feature=cli type=helper control=sequence
//ff:what app.draftBase — a draft file and the published newest version of its DOC-ID (an empty document when never published)
package cli

import (
	"os"

	"github.com/geul-org/geul/go/pkg/core"
)

// draftBase reads the draft at path and, as base, the current version of the
// document its DOC-ID names (resolve.Dir.Current); a document never
// published is compared against an empty one.
func (a *app) draftBase(path string) ([]byte, []byte, error) {
	next, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	issuer, doc, err := docIDOf(next)
	if err != nil {
		return nil, nil, err
	}
	d, err := a.dir()
	if err != nil {
		return nil, nil, err
	}
	base, _, err := d.Current(issuer, doc)
	if base == nil && err == nil {
		base, err = core.EncodeDocument(nil, nil)
	}
	return base, next, err
}
