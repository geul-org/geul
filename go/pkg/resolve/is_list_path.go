//ff:func feature=resolve type=helper control=sequence
//ff:what isListPath — whether a path is a version list document
package resolve

import "strings"

// isListPath reports whether p names a version list document.
func isListPath(p string) bool {
	return strings.HasSuffix(p, "/"+listFile)
}
