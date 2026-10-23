//ff:func feature=toulmin type=encoder control=sequence
//ff:what PutID — append an Edge Id / REF field as an embedded LOCAL Reference record
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// PutID appends v as an embedded LOCAL Reference record (GEUL t002 OVERVIEW:
// Edge Ids and in-document mentions are LOCAL IDs, written in minimal words, so
// no document-wide width is declared). REF 0 ("none") is LOCAL 0. Any uint64
// fits, so it cannot fail.
func PutID(out []core.Word, v TID) []core.Word {
	ws, _ := core.EncodeRecord(core.EncodeLocalRef(uint64(v)))
	return append(out, ws...)
}
