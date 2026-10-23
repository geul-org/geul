//ff:func feature=resolve type=decoder control=sequence
//ff:what isIssuerID — whether s is a canonical 1, 2, 4 or 16-word Issuer ID
package resolve

// isIssuerID reports whether s is a canonical Issuer ID: 4, 8, 16 or 64 hex digits.
func isIssuerID(s string) bool {
	n := (len(s) - 2) / 4
	if n != 1 && n != 2 && n != 4 && n != 16 {
		return false
	}
	_, ok := parseHexID(s, n)
	return ok
}
