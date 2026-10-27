//ff:func feature=notation type=decoder control=selection
//ff:what formatRecord — one embedded record as a token: _, a Reference or Value short form, else its raw words
package notation

import "github.com/geul-org/gogeul/pkg/core"

// formatRecord writes r in the shortest form the package doc lists; a record
// with none is its whole encoding as raw words.
func formatRecord(r core.Record) string {
	if core.IsAbsent(r) {
		return "_"
	}
	var s string
	ok := false
	switch r.Kind {
	case core.KindReference:
		s, ok = formatReference(r)
	case core.KindValue:
		s, ok = formatValue(r)
	}
	if ok {
		return s
	}
	ws, _ := core.EncodeRecord(r)
	return hexToken(ws)
}
