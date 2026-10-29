//ff:func feature=pack type=rule control=sequence
//ff:what isIssuer — whether s is a canonical Issuer ID of 1, 2, 4 or 16 words
package pack

// isIssuer reports whether s is canonical Issuer hex: 4, 8, 16 or 64 digits.
func isIssuer(s string) bool {
	n := len(s) - 2
	if n != 4 && n != 8 && n != 16 && n != 64 {
		return false
	}
	_, ok := parseHex(s, n)
	return ok
}
