//ff:type feature=build type=model
//ff:what baker — one build in progress: the tree root, its source and output .well-known/geul, the codec specs and the count
package build

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// baker walks the archive tree (stored at src) and writes its static copy
// under out, laying documents out by spec.
type baker struct {
	tree  resolve.Store
	src   string
	out   string
	spec  codecspec.Spec
	count Count
}
