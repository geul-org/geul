//ff:func feature=resolve type=helper control=selection
//ff:what codecKey — `proposal/<kind>/<Codec 비트>`, the address of a direct codec's spec document
package resolve

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// codecKey is the address of the spec document of the codec in target
// (GEUL-ORG 직접 코드 코덱북): the Proposal prefix, the Kind and the Codec bits
// as the OVERVIEW tables write them, 4 for a Reference (CCCC) and 7 otherwise.
func codecKey(target core.Word) DocKey {
	k, c := core.SplitTarget(target)
	switch k {
	case core.KindReference:
		return DocKey{Issuer: "proposal/" + kindDirs[k], Doc: fmt.Sprintf("%04b", c>>3)}
	default:
		return DocKey{Issuer: "proposal/" + kindDirs[k], Doc: fmt.Sprintf("%07b", c)}
	}
}
