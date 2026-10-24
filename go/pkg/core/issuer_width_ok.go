//ff:func feature=core type=helper control=selection
//ff:what issuerWidthOK — whether w is an Issuer width (1, 2, 4 or 16 words)
package core

// issuerWidthOK reports whether w words is an Issuer width (GEUL t002 ISSUER).
func issuerWidthOK(w int) bool {
	switch w {
	case 1, 2, 4, 16:
		return true
	default:
		return false
	}
}
