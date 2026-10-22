//ff:func feature=inspect type=helper control=sequence
//ff:what edgeRefs — a LINK's references: src, then dst (the FuncRef ALIAS for do/undo)
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// edgeRefs returns the REF fields of a LINK: src, then dst — the REF, or the
// FuncRef's ALIAS when dst is a function (do/undo).
func edgeRefs(p toulmin.EdgePacket) []uint64 {
	dst := p.DstTID
	if p.DstFunc != nil {
		dst = p.DstFunc.Alias
	}
	return appendRefs(nil, p.Src, dst)
}
