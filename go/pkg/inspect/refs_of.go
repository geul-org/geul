//ff:func feature=inspect type=helper control=selection
//ff:what refsOf — the Edge Ids a decoded TOULMIN packet references, in field order
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// refsOf returns the non-zero Edge Ids data references: a CASE-BEGIN/CASE-END
// its case, a NODE or LINK its REF fields. Other content references nothing.
func refsOf(data any) []uint64 {
	switch p := data.(type) {
	case toulmin.ControlPacket:
		if p.Code != toulmin.ControlCaseBegin && p.Code != toulmin.ControlCaseEnd {
			return nil
		}
		return appendRefs(nil, toulmin.TID(p.Arg))
	case toulmin.NodePacket:
		return nodeRefs(p)
	case toulmin.EdgePacket:
		return edgeRefs(p)
	default:
		return nil
	}
}
