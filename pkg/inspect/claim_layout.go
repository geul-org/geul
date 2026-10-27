//ff:func feature=inspect type=decoder control=sequence
//ff:what claimLayout — the fields of a NODE Claim: subject, predicate, object by kind, mode
package inspect

import "github.com/geul-org/gogeul/pkg/toulmin"

// claimLayout walks a Claim: the subject REF, the predicate word, the object (a
// REF or an UNSIGNED-INTEGER Value, both embedded records, or a SPAN) and the
// mode word.
func claimLayout(w *walker, c toulmin.Claim) {
	w.rec("subject")
	w.words("predicate", 1)
	if c.Kind == toulmin.ObjSpan {
		w.span("object")
	} else {
		w.rec("object")
	}
	w.words("mode", 1)
}
