//ff:func feature=server type=handler control=sequence
//ff:what treeSpec — the codec specs published under the tree's GEUL Issuer, read per request; none when they do not read
package server

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// treeSpec reads the codec spec documents of the tree at root (GEUL-ORG 직접
// 코드 코덱북) on every request, so a republished spec shows at once. A tree
// whose specs do not read gives no spec, and records fall back to the generic
// layout.
func treeSpec(root string) codecspec.Spec {
	s, err := resolve.Dir{Root: root}.CodecSpec()
	if err != nil {
		return nil
	}
	return s
}
