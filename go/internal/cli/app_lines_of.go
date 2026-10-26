//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what app.linesOf — a document's records as notation lines, laid out by the data folder's codec specs
package cli

import (
	"github.com/geul-org/geul/go/pkg/inspect"
	"github.com/geul-org/geul/go/pkg/notation"
)

// linesOf decodes b (inspect.DecodeWith, the data folder's specs) and writes
// each HEAD and BODY record as one line (notation.FormatEntry).
func (a *app) linesOf(b []byte) (docLines, error) {
	doc, err := inspect.DecodeWith(b, a.spec())
	if err != nil {
		return docLines{}, err
	}
	out := docLines{Head: []string{}, Body: []string{}}
	for _, e := range doc.Head {
		out.Head = append(out.Head, notation.FormatEntry(e))
	}
	for _, e := range doc.Body {
		out.Body = append(out.Body, notation.FormatEntry(e))
	}
	return out, nil
}
