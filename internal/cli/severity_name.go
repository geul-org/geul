//ff:func feature=cli type=helper control=selection
//ff:what severityName — a schema severity as error or warn
package cli

import "github.com/geul-org/gogeul/pkg/schema"

// severityName names s.
func severityName(s schema.Severity) string {
	switch s {
	case schema.Error:
		return "error"
	default:
		return "warn"
	}
}
