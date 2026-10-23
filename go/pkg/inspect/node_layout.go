//ff:func feature=inspect type=decoder control=sequence
//ff:what nodeLayout — the fields of a NODE record in flag order
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// nodeLayout walks a NODE record (GEUL t002 TOULMIN NODE): the Edge Id, the
// header word, each flag-gated field in flag order, then the SPAN.
func nodeLayout(w *walker, p toulmin.NodePacket) {
	w.rec("ID")
	w.words("header", 1)
	if p.Func != nil {
		w.rec("func.alias")
		w.words("func.verb", 1)
	}
	if p.Qual != nil {
		w.rec("qualifier")
	}
	if p.Check != nil {
		w.rec("checking")
	}
	if p.Claim != nil {
		claimLayout(w, *p.Claim)
	}
	if p.Verdict != nil {
		w.rec("verdict")
	}
	if p.Flags&toulmin.NodeFlagWith != 0 {
		listLayout(w, "with", len(p.With))
	}
	if p.Flags&toulmin.NodeFlagGrounds != 0 {
		listLayout(w, "grounds", len(p.Grounds))
	}
	w.span("span")
}
