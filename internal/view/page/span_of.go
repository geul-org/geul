//ff:func feature=page type=helper control=selection
//ff:what spanOf — the SOURCE span a decoded NODE, LINK or PROSE record anchors
package page

import "github.com/geul-org/gogeul/pkg/toulmin"

// spanOf returns the SOURCE span of a NODE, a LINK or a PROSE control; false
// for anything without one.
func spanOf(data any) (toulmin.Span, bool) {
	switch p := data.(type) {
	case toulmin.NodePacket:
		return p.Span, true
	case toulmin.EdgePacket:
		return p.Span, true
	case toulmin.ControlPacket:
		if p.Span == nil {
			return toulmin.Span{}, false
		}
		return *p.Span, true
	default:
		return toulmin.Span{}, false
	}
}
