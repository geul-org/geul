//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.latestCodec — the latest live version of a codec spec document, checked against its list hash
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// latestCodec reads the latest live version of target's spec document. A file
// whose content hash differs from its PUBLICATION is a verification failure.
func (d Dir) latestCodec(target core.Word) ([]byte, error) {
	k := codecKey(target)
	lg, err := d.readList(k, core.FormatID(codecDocID(target)))
	if err != nil {
		return nil, err
	}
	e, ok := pickVersion(lg.Entries, 0)
	if !ok {
		return nil, ErrNotFound
	}
	b, err := d.Fetch(context.Background(), versionPath(k, e.Version))
	if err != nil {
		return nil, err
	}
	ws, _ := core.ReadWords(b)
	if h, err := core.ContentHash(ws); err != nil || h != e.Hash {
		return nil, fmt.Errorf("%w: %s version %s", ErrVerify, k.Issuer+"/"+k.Doc, FormatVersion(e.Version))
	}
	return b, nil
}
