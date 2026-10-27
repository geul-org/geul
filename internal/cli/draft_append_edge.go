//ff:func feature=cli type=helper control=sequence
//ff:what draft.appendEdge — add an Edge at the end of the HEAD when it is a HEAD Edge (document SOURCE, QUALIFIER on one), else of the BODY
package cli

import "github.com/geul-org/gogeul/pkg/core"

// appendEdge puts r last in the HEAD when it is a SOURCE on the whole
// document (Target #0) or a QUALIFIER on one of the HEAD's SOURCEs
// (core.HeadEdge, document HEAD와 BODY), and last in the BODY otherwise.
func (d *draft) appendEdge(r core.Record) {
	if _, ok := core.HeadEdge(r, headSourcesOf(d.head)); ok {
		d.head = append(d.head, r)
		return
	}
	d.body = append(d.body, r)
}
