//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what gedcomASCII — every byte is ASCII 0x20~0x7E and the string neither starts nor ends with a space
package schema

// gedcomASCII reports whether b is a GEDCOM date payload's byte form: ASCII
// 0x20~0x7E (a GEDCOM date separates its parts with spaces) with no leading or
// trailing space. b is not empty.
func gedcomASCII(b []byte) bool {
	if b[0] == ' ' || b[len(b)-1] == ' ' {
		return false
	}
	for _, c := range b {
		if c < 0x20 || c > 0x7E {
			return false
		}
	}
	return true
}
