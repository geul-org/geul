//ff:func feature=schema type=helper control=iteration dimension=1
//ff:what edtfPlainDay — whether EDTF bytes are a plain day YYYY-MM-DD (optional leading minus) that GREGORIAN-DATE says exactly
package schema

// edtfPlainDay reports whether b is an EDTF date of the plain form
// [-]YYYY-MM-DD: no qualifier (?, ~, %), no unspecified digit (X), no
// interval, set or Y prefix. Such a day is what GREGORIAN-DATE says exactly,
// so EDTF must not be used for it (V6). Whether the day exists is not checked
// here.
func edtfPlainDay(b []byte) bool {
	s := string(b)
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
	}
	if len(s) != 10 || s[4] != '-' || s[7] != '-' {
		return false
	}
	for i, ch := range s {
		if i != 4 && i != 7 && (ch < '0' || ch > '9') {
			return false
		}
	}
	return true
}
