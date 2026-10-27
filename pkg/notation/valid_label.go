//ff:func feature=notation type=rule control=iteration dimension=1
//ff:what validLabel — whether the text after # is a batch label: ASCII letters, digits, _ and -
package notation

// validLabel reports whether s, the text of a `#label`, is not empty and
// holds only ASCII letters, digits, `_` and `-`.
func validLabel(s string) bool {
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' || c == '-') {
			return false
		}
	}
	return s != ""
}
