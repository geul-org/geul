//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.publishAt — give a document the next version at an address and append it to that address's version list
package resolve

import (
	"errors"
	"time"

	"github.com/geul-org/geul/go/pkg/core"
)

// publishAt stores b (content hash h) at k as the next version of document
// doc of Issuer issuer (the PUBLICATION's DOC-POINTER): `<k>/<버전 ID>.geul`, `h/<내용 해시>.geul`, and a new version list
// `<k>/versions.geul` that keeps the previous records, appends a PUBLICATION
// and points PARENT at the previous list, which stays under `h/`. Content
// already in the list returns that entry and writes nothing.
func (d Dir) publishAt(k DocKey, issuer, doc []core.Word, b []byte, h [32]byte) (VersionEntry, error) {
	lg, err := d.readList(k, core.FormatID(doc))
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
	rec := core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: uint64(len(lg.Records)) + 1, Issuer: issuer, Doc: doc, Version: v, Hash: h})
	var parent *[32]byte
	if lg.Present {
		parent = &lg.Hash
	}
	list := encodeVersionList(append(lg.Records, rec), parent, time.Now().Unix(), nextLocalID(lg.Next, b))
	if err := d.writeAll(map[string][]byte{versionPath(k, v): b, hashPath(h): b, hashPath(listHash(list)): list, listPath(k): list}); err != nil {
		return VersionEntry{}, err
	}
	return VersionEntry{Version: v, Hash: h}, nil
}
