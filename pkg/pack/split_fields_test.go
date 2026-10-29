//ff:func feature=pack type=helper control=sequence
//ff:what splitFields — strings.Fields for test tables
package pack

import "strings"

// splitFields splits s at spaces.
func splitFields(s string) []string {
	return strings.Fields(s)
}
