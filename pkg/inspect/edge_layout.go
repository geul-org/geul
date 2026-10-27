//ff:func feature=inspect type=decoder control=sequence
//ff:what edgeLayout — the fields of a LINK record: Edge Id, src, dst, flags, gate, SPAN
package inspect

import "github.com/geul-org/gogeul/pkg/toulmin"

// edgeLayout walks a LINK record (GEUL t002 TOULMIN LINK): the Edge Id, the src
// REF, the dst (a FuncRef for do/undo, else a REF), the flag word, the gate
// operator and percent Value when gated, then the SPAN.
func edgeLayout(w *walker, p toulmin.EdgePacket) {
	w.rec("ID")
	w.rec("src")
	if p.DstFunc != nil {
		w.rec("dst.alias")
		w.words("dst.verb", 1)
	} else {
		w.rec("dst")
	}
	w.words("flags", 1)
	if p.Gate != nil {
		w.words("gate.op", 1)
		w.rec("gate")
	}
	w.span("span")
}
