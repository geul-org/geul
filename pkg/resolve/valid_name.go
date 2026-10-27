//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what validName — whether s is a document name: a-z, 0-9 and '-', not at either end, never starting with "0x"
package resolve

import "strings"

// validName follows GEUL-ORG 문서 이름 so a name never looks like an ID.
func validName(s string) bool {
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
