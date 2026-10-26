//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what docIDOf — the Issuer and document ID a document's DOC-ID names
package cli

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// docIDOf reads b's HEAD for DOC-ID. A document without one is an error:
// only a document with a DOC-ID has versions.
func docIDOf(b []byte) ([]core.Word, []core.Word, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, nil, err
	}
	for _, r := range doc.Head {
		if issuer, id, err := core.DecodeDocID(r); err == nil {
			return issuer, id, nil
		}
	}
	return nil, nil, errors.New("the document has no DOC-ID")
}
