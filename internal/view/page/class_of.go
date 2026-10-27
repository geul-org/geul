//ff:func feature=page type=helper control=selection
//ff:what classOf — the style class of a record from its decoded content
package page

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// classOf names the style class of decoded content; "raw" when the record has
// no typed reading.
func classOf(data any) string {
	switch p := data.(type) {
	case inspect.Value:
		return "value"
	case inspect.Source:
		return "source"
	case toulmin.TextPacket:
		return "text"
	case toulmin.ControlPacket:
		if p.Code == toulmin.ControlSection {
			return "section"
		}
		return "control"
	case toulmin.NodePacket:
		return "node"
	case toulmin.EdgePacket:
		return "link"
	default:
		return "raw"
	}
}
