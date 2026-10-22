//ff:func feature=inspect type=helper control=selection
//ff:what idOf — the Edge Id a decoded TOULMIN packet carries; 0 for other content
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// idOf returns the Edge Id of a decoded TOULMIN record, 0 for anything else.
func idOf(data any) uint64 {
	switch p := data.(type) {
	case Source:
		return uint64(p.TID)
	case toulmin.TextPacket:
		return uint64(p.TID)
	case toulmin.ControlPacket:
		return uint64(p.ID)
	case toulmin.NodePacket:
		return uint64(p.TID)
	case toulmin.EdgePacket:
		return uint64(p.ID)
	default:
		return 0
	}
}
