//ff:func feature=cli type=helper control=sequence
//ff:what app.nextFree — the next LOCAL ID a draft may assign: above its own IDs and its published NEXT-LOCAL-ID
package cli

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// nextFree returns max(NEXT-LOCAL-ID of the draft's published version list,
// the draft's highest LOCAL ID + 1) (README 작성 명령: a draft is not in the
// version list yet). Deleted IDs stay unused because the list's counter never
// goes down.
func (a *app) nextFree(d draft) uint64 {
	b, _ := core.EncodeDocument(d.head, d.body)
	doc, _ := inspect.Decode(b)
	next := inspect.MaxLocalID(doc) + 1
	issuer, id, err := docIDOf(b)
	dir, derr := a.dir()
	if err != nil || derr != nil {
		return next
	}
	if h, err := dir.Log(issuer, id); err == nil && h.Next > next {
		next = h.Next
	}
	return next
}
