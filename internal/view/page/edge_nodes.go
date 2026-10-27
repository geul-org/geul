//ff:func feature=page type=helper control=iteration dimension=1
//ff:what edgeNodes — the node records inside one Edge: embedded non-Value records other than its Edge Id
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// edgeNodes returns the embedded records of Edge e that are nodes: payload
// records other than its Edge Id that are not Values. A LOCAL naming another
// Edge is a Reference too and counts as a node.
func edgeNodes(e inspect.Entry) []inspect.Entry {
	if e.Record.Kind != core.KindEdge {
		return nil
	}
	var out []inspect.Entry
	for _, p := range inspect.Layout(e) {
		if p.Head || !p.Record || p.Name == "ID" {
			continue
		}
		r, _, _ := core.DecodeRecord(p.Words)
		if r.Kind == core.KindValue {
			continue
		}
		out = append(out, inspect.Entry{Record: r, Words: p.Words, Name: core.CodecName(r.Kind, r.Codec)})
	}
	return out
}
