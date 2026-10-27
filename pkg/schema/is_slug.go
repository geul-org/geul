//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what isSlug — a document name: lowercase ASCII letters, digits and '-', not at the ends, not starting with 0x
package schema

import "strings"

// isSlug reports whether s is a valid document name (GEUL-ORG 문서 이름):
// a-z, 0-9 and '-', no '-' at either end, and no "0x" prefix so it never looks
// like an ID.
func isSlug(s string) bool {
	if s == "" || s[0] == '-' || s[len(s)-1] == '-' || strings.HasPrefix(s, "0x") {
		return false
	}
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-') {
			return false
		}
	}
	return true
}
