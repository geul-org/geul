//ff:func feature=cli type=helper control=sequence
//ff:what draft.edgeAt — where the Edge with an Edge Id is: the BODY or the HEAD and its position; -1 when none
package cli

import "github.com/geul-org/gogeul/pkg/core"

// edgeAt finds the Edge whose Edge Id is id in the BODY, then in the HEAD
// (a document SOURCE or its QUALIFIER), and returns that part of d and the
// position (edgeIndex); the position is -1 when neither has it.
func (d *draft) edgeAt(id uint64) (*[]core.Record, int) {
	if i := edgeIndex(d.body, id); i >= 0 {
		return &d.body, i
	}
	return &d.head, edgeIndex(d.head, id)
}
