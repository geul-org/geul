//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what lineOf — the notation line of a record as it appears in a decoded document; its raw words when not found
package cli

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
	"github.com/geul-org/geul/go/pkg/notation"
)

// lineOf finds the entry of doc (HEAD, then BODY) encoding to r's words and
// formats it (notation.FormatEntry); r alone is used when none matches.
func lineOf(doc inspect.Document, r core.Record) string {
	ws, _ := core.EncodeRecord(r)
	for _, e := range append(append([]inspect.Entry{}, doc.Head...), doc.Body...) {
		if slices.Equal(e.Words, ws) {
			return notation.FormatEntry(e)
		}
	}
	return notation.FormatEntry(inspect.Entry{Record: r, Words: ws})
}
