//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what decodeHead — HEAD entries, after a first pass that reads the declared widths
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// decodeHead interprets each HEAD Meta record in order and returns the
// declared widths and PROFILE text. ID-WIDTH is read first because an
// IMPORT-DOC before it still uses its IMPORT ID width.
func decodeHead(head []core.Record) ([]Entry, Widths, string) {
	w := Widths{Edge: 1, Var: 1, Import: 1}
	for _, r := range head {
		if pw, ok := parseWidths(r); ok && r.Codec == core.MetaIDWidth {
			w = pw
		}
	}
	entries := make([]Entry, 0, len(head))
	profile := ""
	for _, r := range head {
		e := headEntry(r, w)
		if v, ok := e.Data.(Value); ok && r.Codec == core.MetaProfile {
			profile = v.Text
		}
		entries = append(entries, e)
	}
	return entries, w, profile
}
