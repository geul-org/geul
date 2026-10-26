//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what referrer — the first Edge of a draft that refers to one of the given Edge Ids
package cli

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// referrer decodes d and returns the Edge Id of the first BODY entry that
// mentions an ID of ids (inspect.Mentions).
func referrer(d draft, ids []uint64) (uint64, bool) {
	b, _ := core.EncodeDocument(d.head, d.body)
	doc, _ := inspect.Decode(b)
	for _, e := range doc.Body {
		if slices.ContainsFunc(inspect.Mentions(e), func(r uint64) bool { return slices.Contains(ids, r) }) {
			return e.ID, true
		}
	}
	return 0, false
}
