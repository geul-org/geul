//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.PublishCodec — publish a direct codec's spec document at proposal/<kind>/<Codec 비트>/
package resolve

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// PublishCodec stores the codec spec document b (edge/DECLARE 규격 문서) at
// the address of the codec its DECLARE targets, with versions and a version
// list like an Issuer's document. The list's document ID is the target word
// widened to 4 words. b must hold exactly one DECLARE, of a direct codec.
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
	return d.publishAt(codecKey(target), []core.Word{0, 0, 0, target}, b, h)
}
