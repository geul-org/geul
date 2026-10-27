//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what decodeHead — HEAD entries in order plus the PROFILE text
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// decodeHead interprets each HEAD Meta record in order and returns the entries
// and the PROFILE text ("" when absent). Every field width follows from the
// codec and Length, so no record has to be read ahead of another.
func decodeHead(head []core.Record) ([]Entry, string) {
	entries := make([]Entry, 0, len(head))
	profile := ""
	for _, r := range head {
		e := headEntry(r)
		if v, ok := e.Data.(Value); ok && r.Codec == core.MetaProfile {
			profile = v.Text
		}
		entries = append(entries, e)
	}
	return entries, profile
}
