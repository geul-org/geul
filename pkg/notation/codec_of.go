//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what codecOf — the direct codec of a Kind whose core spec name is name (case ignored); Reference names give CCCC with Quantifier 0
package notation

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// codecOf looks name up among the core codec names of kind (core.CodecName).
// ISSUER codecs, which need an Issuer, are not direct and are skipped.
func codecOf(kind core.Kind, name string) (uint8, bool) {
	if name == "" {
		return 0, false
	}
	up := strings.ToUpper(name)
	for c := 0; c < 128; c++ {
		codec := uint8(c)
		if kind == core.KindReference {
			codec = uint8(c) << 3
		}
		if core.IssuerWidth(kind, codec) == 0 && core.CodecName(kind, codec) == up && (kind != core.KindReference || c < 16) {
			return codec, true
		}
	}
	return 0, false
}
