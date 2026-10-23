//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Publish — assign the next version ID to a named document and store it with its list entry
package resolve

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// Publish stores the document b in the local tree as the publisher does
// (GEUL-ORG 공표 주소): it names the document by its DOC-ID, gives it the next
// version ID, writes `<Issuer>/<문서 ID>/<버전 ID>.geul` and `h/<내용 해시>.geul`,
// and appends the entry to the version list. Publishing content already in the
// list returns that entry and writes nothing.
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
	name, ok := docIDOf(doc)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: publishing needs a DOC-ID")
	}
	k := Key{Issuer: name.Issuer, Doc: name.Doc}.DocKey()
	list, err := d.readList(k)
	if err != nil {
		return VersionEntry{}, err
	}
	if e, ok := findHash(list, h); ok {
		return e, nil
	}
	v, ok := nextVersion(list)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: version IDs used up; publish under a new document ID")
	}
	e := VersionEntry{Version: v, Hash: h}
	if err := d.write(versionPath(k, v), b); err != nil {
		return VersionEntry{}, err
	}
	if err := d.write(hashPath(h), b); err != nil {
		return VersionEntry{}, err
	}
	return e, d.write(listPath(k), []byte(formatVersionList(append(list, e))))
}
