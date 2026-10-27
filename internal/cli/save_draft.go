//ff:func feature=cli type=helper control=sequence
//ff:what saveDraft — write a working file only when the result keeps the spec
package cli

import (
	"os"

	"github.com/geul-org/gogeul/pkg/core"
)

// saveDraft encodes d and writes it to path, unless it has schema errors
// (schemaErrors): an edit that breaks the document is not written. Its
// records were decoded or read from notation, so they encode.
func saveDraft(path string, d draft) error {
	b, _ := core.EncodeDocument(d.head, d.body)
	if err := schemaErrors(b); err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}
