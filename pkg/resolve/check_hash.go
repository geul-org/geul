//ff:func feature=resolve type=resolver control=selection
//ff:what checkHash — a document returned only when its content hash is the expected one
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkHash returns b when its content hash (core.ContentHash) is want, else
// an ErrVerify error.
func checkHash(b []byte, want [32]byte) ([]byte, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	h, err := core.ContentHash(ws)
	switch {
	case err != nil:
		return nil, fmt.Errorf("%w: %v", ErrVerify, err)
	case h != want:
		return nil, fmt.Errorf("%w: content hash differs from the version list", ErrVerify)
	default:
		return b, nil
	}
}
