//ff:func feature=geulmd type=helper control=selection
//ff:what kindName — the GEULMD keyword of a record kind
package geulmd

import "github.com/geul-org/geul/go/pkg/core"

// kindName returns the keyword of k used by raw RECORD lines.
func kindName(k core.Kind) string {
	switch k {
	case core.KindReference:
		return "REFERENCE"
	case core.KindValue:
		return "VALUE"
	case core.KindEdge:
		return "EDGE"
	default:
		return "META"
	}
}
