//ff:func feature=cli type=helper control=sequence
//ff:what readDraft — a working file split into HEAD and BODY records
package cli

import (
	"os"

	"github.com/geul-org/gogeul/pkg/core"
)

// readDraft reads the document at path.
func readDraft(path string) (draft, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return draft{}, err
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return draft{}, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return draft{}, err
	}
	return draft{head: doc.Head, body: doc.Body}, nil
}
