//ff:func feature=inspect type=decoder control=selection
//ff:what typedLayout — walk the payload field by field for content whose layout is known; false otherwise
package inspect

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// typedLayout walks w through the fields of a decoded version list record
// (edge/PUBLICATION) or TOULMIN record (GEUL t002 TOULMIN): SOURCE, TEXT,
// framing controls, NODE and LINK. It reports false for any other content.
func typedLayout(w *walker, data any) bool {
	switch p := data.(type) {
	case core.Publication:
		publicationLayout(w, p)
	case Source:
		w.rec("ID")
		w.words("span width", 1)
		w.rec("text")
	case toulmin.TextPacket:
		w.rec("ID")
		w.rec("text")
	case toulmin.ControlPacket:
		controlLayout(w, p)
	case toulmin.NodePacket:
		nodeLayout(w, p)
	case toulmin.EdgePacket:
		edgeLayout(w, p)
	default:
		return false
	}
	return true
}
