//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what Mentions — the LOCAL IDs a record mentions among its laid-out record parts, an Edge's own Edge Id left out
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// Mentions returns, in order, every LOCAL Reference among e's payload record
// parts (Layout) except an Edge's leading Edge Id, which declares rather than
// mentions (OVERVIEW Reference). Unlike Entry.Refs it needs no typed content,
// so it also covers codecs laid out generically. A LOCAL inside a Value's own
// payload is not looked into.
func Mentions(e Entry) []uint64 {
	var out []uint64
	first := e.Record.Kind == core.KindEdge
	for _, p := range Layout(e) {
		if p.Head || !p.Record {
			continue
		}
		r, _, _ := core.DecodeRecord(p.Words)
		id, err := core.DecodeLocalRef(r)
		if err == nil && !first {
			out = append(out, id)
		}
		first = false
	}
	return out
}
