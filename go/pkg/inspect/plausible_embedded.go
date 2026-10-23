//ff:func feature=inspect type=decoder control=selection
//ff:what plausibleEmbedded — whether a decoded record looks like an embedded Reference or Value
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// plausibleEmbedded accepts a Reference that carries a target ID and a Value
// whose codec this library names or that is an ISSUER codec.
func plausibleEmbedded(r core.Record) bool {
	switch r.Kind {
	case core.KindReference:
		return len(r.Payload) > 0 && core.CodecName(r.Kind, r.Codec) != ""
	case core.KindValue:
		return core.CodecName(r.Kind, r.Codec) != ""
	default:
		return false
	}
}
