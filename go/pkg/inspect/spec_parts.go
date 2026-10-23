//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what specParts — a payload cut by its codec spec: Edge Id, then each declared field; false when it does not fit
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// specParts lays out e's payload by e.Spec (edge/DECLARE 해석 규칙): an Edge's
// leading Edge Id, then each field in order, a raw field as its fixed words,
// a record field by its own Length and the rest field as what the fixed
// fields after it leave (tailWidth). It is false when there is no spec or
// the fields do not use the payload exactly.
func specParts(e Entry) ([]Part, bool) {
	if e.Spec == nil {
		return nil, false
	}
	tail, ok := tailWidth(e.Spec)
	if !ok {
		return nil, false
	}
	w := &walker{ws: e.Record.Payload, tail: tail}
	if e.Record.Kind == core.KindEdge {
		w.rec("ID")
	}
	for _, f := range e.Spec {
		w.field(f)
	}
	return w.out, !w.bad && w.pos == len(w.ws)
}
