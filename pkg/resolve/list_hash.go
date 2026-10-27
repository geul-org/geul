//ff:func feature=resolve type=helper control=sequence
//ff:what listHash — the content hash of a version list document this package encoded
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// listHash is the content hash of a list document from encodeVersionList,
// which is always well formed.
func listHash(list []byte) [32]byte {
	ws, _ := core.ReadWords(list)
	h, _ := core.ContentHash(ws)
	return h
}
