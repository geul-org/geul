//ff:func feature=inspect type=decoder control=selection
//ff:what controlLayout — the fields of a framing control: Edge Id, then a SPAN, a section code or a case REF
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// controlLayout walks a SECTION, CASE-BEGIN/END or PROSE record: the Edge Id,
// then the PROSE span, the section code word or the case REF.
func controlLayout(w *walker, p toulmin.ControlPacket) {
	w.rec("ID")
	switch p.Code {
	case toulmin.ControlProse:
		w.span("span")
	case toulmin.ControlSection:
		w.words("section", 1)
	default:
		w.rec("case")
	}
}
