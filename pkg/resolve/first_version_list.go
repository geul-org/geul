//ff:func feature=resolve type=resolver control=sequence
//ff:what FirstVersionList — the version list Overwrite writes for a document published as version 1 with no earlier list, recorded at a given time
package resolve

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// FirstVersionList returns the bytes of `versions.geul` for document b
// published as its only version 1 with no earlier list (edge/PUBLICATION),
// exactly as Overwrite and a first Publish write it but with RECORDED-AT at
// (Unix seconds) instead of the clock: one PUBLICATION of b's content hash,
// no PARENT, NEXT-LOCAL-ID one above b's highest LOCAL ID. Writers that store
// archives elsewhere (a pack stream) use it to give the same bytes; a fixed
// at keeps a run reproducible.
func FirstVersionList(b []byte, at int64) ([]byte, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return nil, err
	}
	doc, _ := inspect.Decode(b)
	id, ok := docIDOf(doc)
	if !ok {
		return nil, errors.New("geul/resolve: publishing needs a DOC-ID")
	}
	rec := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: id.Issuer, Doc: id.Doc, Version: 1, Hash: h})
	return encodeVersionList([]core.Record{rec}, nil, at, nextLocalID(0, b)), nil
}
