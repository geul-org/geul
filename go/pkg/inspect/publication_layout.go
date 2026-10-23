//ff:func feature=inspect type=decoder control=selection
//ff:what publicationLayout — walk a PUBLICATION, DEPRECATION or TOMBSTONE payload: Edge Id, document, version, then hash or reason
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// publicationLayout walks the fields of a version list record (GEUL t002
// docs/edge/PUBLICATION): the LOCAL Edge Id, the 4-word document ID, the 2-word
// version ID, then the content hash (PUBLICATION) or the reason (TOMBSTONE).
func publicationLayout(w *walker, p core.Publication) {
	w.rec("ID")
	w.words("doc", 4)
	w.words("version", 2)
	switch p.Codec {
	case core.EdgePublication:
		w.words("hash", 16)
	case core.EdgeTombstone:
		w.words("reason", 1)
	}
}
