//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.PublishCodec — publish a direct codec's spec document at proposal/<kind>/<Codec 비트>/
package resolve

import (
	"errors"
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// PublishCodec stores the codec spec document b (edge/DECLARE 규격 문서) at
// the address of the codec its DECLARE targets, with versions and a version
// list like an Issuer's document. A spec document is a document of the GEUL
// Issuer 0x0000 whose document ID is codecDocID of its target word; b
// must hold exactly one DECLARE, of a direct codec, and its DOC-ID must name
// that document.
func (d Dir) PublishCodec(b []byte) (VersionEntry, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return VersionEntry{}, err
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return VersionEntry{}, err
	}
	doc, _ := core.SplitDocument(ws)
	target, ok := declareTarget(doc.Body)
	if !ok {
		return VersionEntry{}, errors.New("geul/resolve: a codec spec document needs exactly one DECLARE of a direct codec")
	}
	issuer, docID := []core.Word{0}, codecDocID(target)
	parsed, _ := inspect.Decode(b)
	if id, ok := docIDOf(parsed); !ok || !slices.Equal(id.Issuer, issuer) || !slices.Equal(id.Doc, docID) {
		return VersionEntry{}, errors.New("geul/resolve: a codec spec document needs DOC-ID DOC-POINTER(0x0000, 0x000000000001<codec TAG>, 0)")
	}
	return d.publishAt(codecKey(target), issuer, docID, b, h)
}
