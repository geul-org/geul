//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what parseCodecKey — the target word of a `proposal/<kind>/<Codec 비트>` address
package resolve

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseCodecKey reads a codec spec address back into its target word: the Kind
// segment and the Codec bits, 4 for a Reference and 7 otherwise. It is false
// for an unknown Kind or bits of the wrong width.
func parseCodecKey(kind, bits string) (core.Word, bool) {
	for k, name := range kindDirs {
		c, err := strconv.ParseUint(bits, 2, 8)
		if name != kind || err != nil {
			continue
		}
		t := core.TargetWord(k, uint8(c))
		if k == core.KindReference {
			t = core.TargetWord(k, uint8(c)<<3)
		}
		return t, codecKey(t).Doc == bits
	}
	return 0, false
}
