//ff:func feature=page type=builder control=selection
//ff:what bodyFields — the fields of a BODY record from its decoded content
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// bodyFields lays out decoded BODY content: a Value, the SOURCE's size and SPAN
// width, a TEXT's string, or the fields of a control, NODE or LINK. A record
// with no typed reading shows its codec spec fields, else its Issuer and format, if any.
func bodyFields(e inspect.Entry, labels map[uint64]string) []Field {
	switch p := e.Data.(type) {
	case inspect.Value:
		return []Field{{Key: "value", Text: valueText(p)}}
	case inspect.Source:
		return []Field{{Key: "bytes", Text: strconv.Itoa(len(p.Bytes))}, {Key: "span width", Text: strconv.Itoa(p.SpanWidth)}}
	case toulmin.TextPacket:
		return []Field{{Key: "text", Text: strconv.Quote(string(p.Bytes))}}
	case toulmin.ControlPacket:
		return controlFields(p, labels)
	case toulmin.NodePacket:
		return nodeFields(p, labels)
	case toulmin.EdgePacket:
		return linkFields(p, labels)
	default:
		return specFields(e)
	}
}
