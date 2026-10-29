//ff:func feature=pack type=helper control=sequence
//ff:what contentHash — the GEUL content hash of document bytes (core.ContentHash), false when they are not one document
package pack

import "github.com/geul-org/gogeul/pkg/core"

// contentHash computes the `h/` name of b.
func contentHash(b []byte) ([32]byte, bool) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return [32]byte{}, false
	}
	h, err := core.ContentHash(ws)
	return h, err == nil
}
