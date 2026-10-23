//ff:func feature=inspect type=decoder control=selection
//ff:what genericParts — a payload without a typed layout: a leading embedded record when the spec fixes one, then the rest
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// genericParts splits a payload whose fields are not known. An Edge payload
// opens with its LOCAL Edge Id (GEUL t002 OVERVIEW Edge) and a Meta Value
// payload is one embedded Value, so a leading record that decodes is its own
// part; whatever follows is one "payload" part.
func genericParts(e Entry) []Part {
	w := &walker{ws: e.Record.Payload}
	_, isValue := e.Data.(Value)
	r, _, err := core.DecodeRecord(w.ws)
	switch {
	case err == nil && e.Record.Kind == core.KindEdge && r.Kind == core.KindReference && r.Codec == core.CodecLocalIdentified:
		w.rec("ID")
	case err == nil && e.Record.Kind == core.KindMeta && isValue:
		w.rec("value")
	}
	if w.pos < len(w.ws) {
		w.words("payload", len(w.ws)-w.pos)
	}
	return w.out
}
