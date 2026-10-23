//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Publish — assign the next version ID to a document and append it to its version list document
package resolve

import (
	"errors"
	"time"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// Publish stores the document b in the local tree as the publisher does
// (GEUL-ORG 공표 주소, edge/PUBLICATION): it identifies the document by its
// DOC-ID, gives it the next version ID, writes `<Issuer>/<문서 ID>/<버전 ID>.geul`
// and `h/<내용 해시>.geul`, and writes a new version list document
// (`<Issuer>/<문서 ID>/versions.geul`) that keeps the previous records, appends
// a PUBLICATION and points PARENT at the previous list, which stays under `h/`.
// Publishing content already in the list returns that entry and writes nothing.
func (d Dir) Publish(b []byte) (VersionEntry, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	doc, _ := inspect.Decode(b)
	id, ok := docIDOf(doc)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: publishing needs a DOC-ID")
	}
	k := Key{Issuer: id.Issuer, Doc: id.Doc}.DocKey()
	lg, err := d.readList(k)
	if err != nil {
		return VersionEntry{}, err
	}
	if e, ok := findHash(lg.Entries, h); ok {
		return e, nil
	}
	v, ok := nextVersion(lg.Entries)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: version IDs used up; publish under a new document ID")
	}
	rec := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: uint64(len(lg.Records)) + 1, Doc: id.Doc, Version: v, Hash: h})
	var parent *[32]byte
	if lg.Present {
		parent = &lg.Hash
	}
	list := encodeVersionList(append(lg.Records, rec), parent, time.Now().Unix())
	if err := d.writeAll(map[string][]byte{versionPath(k, v): b, hashPath(h): b, hashPath(listHash(list)): list, listPath(k): list}); err != nil {
		return VersionEntry{}, err
	}
	return VersionEntry{Version: v, Hash: h}, nil
}
