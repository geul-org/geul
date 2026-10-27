//ff:func feature=resolve type=rule control=sequence
//ff:what verifyDoc — check fetched bytes against the list's content hash and the Key's DOC-ID
package resolve

import (
	"fmt"
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// verifyDoc accepts b as version v of k only when its content hash is the one
// the version list gives and its DOC-ID is k's Issuer and document ID.
func verifyDoc(b []byte, k Key, v VersionEntry) (*Node, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	h, err := core.ContentHash(ws)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	if h != v.Hash {
		return nil, fmt.Errorf("%w: content hash of version %s", ErrVerify, FormatVersion(v.Version))
	}
	doc, _ := inspect.Decode(b)
	name, ok := docIDOf(doc)
	if !ok || !slices.Equal(name.Issuer, k.Issuer) || !slices.Equal(name.Doc, k.Doc) {
		return nil, fmt.Errorf("%w: DOC-ID is not %s/%s", ErrVerify, core.FormatID(k.Issuer), core.FormatID(k.Doc))
	}
	return &Node{Key: k, Version: v.Version, Hash: h, Doc: doc}, nil
}
