//ff:func feature=cli type=helper control=selection
//ff:what stateOf — a version's state as a suffix: deleted, deprecated or nothing
package cli

import "github.com/geul-org/geul/go/pkg/resolve"

// stateOf is "\tdeleted", "\tdeprecated" or "".
func stateOf(e resolve.VersionEntry) string {
	switch {
	case e.Deleted:
		return "\tdeleted"
	case e.Deprecated:
		return "\tdeprecated"
	default:
		return ""
	}
}
